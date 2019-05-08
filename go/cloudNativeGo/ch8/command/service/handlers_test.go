package service

import (
	"fmt"
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch8/command/fakes"
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch8/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

var engine *gin.Engine = nil
var handler *Handlers = nil
var dispatch *fakes.FakeQueueDispatcher = nil

func init() {
	engine := getEngine()
	dispatch = fakes.NewFakeQueueDispatcher()
	handler = initHandler(dispatch)
	initRouter(engine, handler)
}

func TestHandlers_AddTelemetryCommand(t *testing.T) {
	t.Run("return 400 when post empty form", func(t *testing.T) {
		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/command/telemetry", nil)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		engine.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("return 400 when post form without drone_id", func(t *testing.T) {
		form := url.Values{}
		form.Set("battery", "90")
		form.Set("uptime", "20")
		form.Set("core_temp", "50")
		body := strings.NewReader(form.Encode())

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/command/telemetry", body)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		engine.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("return 400 when post form without battery", func(t *testing.T) {
		form := url.Values{}
		id := common.NewDroneId()
		form.Set("drone_id", strconv.FormatInt(id, 10))
		form.Set("uptime", "20")
		form.Set("core_temp", "50")
		body := strings.NewReader(form.Encode())

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/command/telemetry", body)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		engine.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("return 400 when post form without uptime", func(t *testing.T) {
		form := url.Values{}
		id := common.NewDroneId()
		form.Set("drone_id", strconv.FormatInt(id, 10))
		form.Set("battery", "80")
		form.Set("core_temp", "50")
		body := strings.NewReader(form.Encode())

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/command/telemetry", body)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		engine.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("return 400 when post form without core_temp", func(t *testing.T) {
		form := url.Values{}
		id := common.NewDroneId()
		form.Set("drone_id", strconv.FormatInt(id, 10))
		form.Set("battery", "80")
		form.Set("uptime", "20")
		body := strings.NewReader(form.Encode())

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/command/telemetry", body)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		engine.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("return 201 and send command when post verified", func(t *testing.T) {
		form := url.Values{}
		id := common.NewDroneId()
		form.Set("drone_id", strconv.FormatInt(id, 10))
		form.Set("battery", "80")
		form.Set("uptime", "20")
		form.Set("core_temp", "50")
		body := strings.NewReader(form.Encode())

		command := &telemetryCommand{
			DroneID:id,
			RemainingBattery:80,
			Uptime:20,
			CoreTemp:50,
		}

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/command/telemetry", body)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		engine.ServeHTTP(res, req)

		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, command, dispatch.Messages[len(dispatch.Messages)-1])
	})
}

func TestHandlers_AddAlertCommand(t *testing.T) {
	id := common.NewDroneId()
	cases := []struct{
		name string
		command alertCommand
		wantCode int
	} {
		{"return 400 when post empty form", alertCommand{}, 400},
		{"return 400 when post without drone_id", alertCommand{0, 001, "low battery"}, 400},
		{"return 400 when post without fault_code", alertCommand{id, 0, "low battery"}, 400},
		{"return 400 when post without description", alertCommand{id, 001, ""}, 400},
		{"return 200 and send command when post form verified", alertCommand{id, 001, "low battery"}, 201},
	}

	for _,c := range cases {
		t.Run(c.name, func(t *testing.T) {
			form := url.Values{}
			if c.command.DroneID == 0 {
				form.Set("drone_id", "")
			} else {
				form.Set("drone_id", strconv.FormatInt(id, 10))
			}
			form.Set("fault_code", strconv.Itoa(c.command.FaultCode))
			form.Set("description", c.command.Description)
			body := strings.NewReader(form.Encode())

			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/command/alerts", body)
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			engine.ServeHTTP(res, req)

			assert.Equal(t, c.wantCode, res.Code)
			if c.wantCode == 201 {
				assert.Equal(t, &c.command, dispatch.Messages[len(dispatch.Messages)-1])
			}
		})
	}
}

func TestHandlers_AddPositionCommand(t *testing.T) {
	id := common.NewDroneId()
	cases := []struct{
		name string
		command positionCommand
		wantCode int
	} {
		{"return 400 when post empty form", positionCommand{}, 400},
		{"return 400 when post without drone_id", positionCommand{0, 30.8,  60.2, 100.1, 200.5, 300}, 400},
		{"return 400 when post without latitude", positionCommand{id, 0,  60.2, 100.1, 200.5, 300}, 400},
		{"return 400 when post without longitude", positionCommand{id, 30.8,  0, 100.1, 200.5, 300}, 400},
		{"return 400 when post without altitude", positionCommand{id, 30.8,  60.2, 0, 200.5, 300}, 400},
		{"return 400 when post without speed", positionCommand{id, 30.8,  60.2, 100.1, 0, 300}, 400},
		{"return 400 when post without heading_cardinal", positionCommand{id, 30.8,  60.2, 100.1, 200.5, 0}, 400},
		{"return 201 and send command when post form verified", positionCommand{id, 30.8,  60.2, 100.1, 200.5, 300}, 201},
	}

	for _,c := range cases {
		t.Run(c.name, func(t *testing.T) {
			form := url.Values{}
			if c.command.DroneID == 0 {
				form.Set("drone_id", "")
			} else {
				form.Set("drone_id", strconv.FormatInt(id, 10))
			}
			form.Set("latitude", fmt.Sprintf("%f", c.command.Latitude))
			form.Set("longitude", fmt.Sprintf("%f", c.command.Longitude))
			form.Set("altitude", fmt.Sprintf("%f", c.command.Altitude))
			if c.command.CurrentSpeed == 0 {
				form.Set("current_speed", "")
			} else {
				form.Set("current_speed", fmt.Sprintf("%f", c.command.CurrentSpeed))
			}
			form.Set("current_speed", fmt.Sprintf("%f", c.command.CurrentSpeed))
			form.Set("heading_cardinal", fmt.Sprintf("%d", c.command.HeadingCardinal))
			body := strings.NewReader(form.Encode())

			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/command/positions", body)
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			engine.ServeHTTP(res, req)

			assert.Equal(t, c.wantCode, res.Code)
			if c.wantCode == 201 {
				assert.Equal(t, &c.command, dispatch.Messages[len(dispatch.Messages)-1])
			}
		})
	}
}

func getEngine() *gin.Engine {
	if engine == nil {
		engine = newEngine()
	}

	return engine
}

func newEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine = gin.Default()

	return engine
}

func initHandler(dispatch queueDispatcher) *Handlers {
	if handler != nil {
		return handler
	}
	handler = &Handlers{dispatch:dispatch}

	return handler
}

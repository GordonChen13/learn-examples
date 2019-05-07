package service

import (
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch8/command/fakes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var engine *gin.Engine = nil
var handler *Handlers = nil

func TestHandlers_AddTelemetryCommand(t *testing.T) {
	engine := getEngine()
	handler = initHandler()
	initRouter(engine, handler)

	t.Run("return 400 when post empty form", func(t *testing.T) {
		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/command/telemetry", nil)
		engine.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("return 400 when post form without stock", func(t *testing.T) {
		form := url.Values{}
		form.Set("sku", "XiaoMi 9")
		form.Set("shipswithin", "20")
		body := strings.NewReader(form.Encode())

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/command/telemetry", body)
		engine.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

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

func initHandler() *Handlers {
	if handler != nil {
		return handler
	}
	dispatch := fakes.NewFakeQueueDispatcher()
	handler = &Handlers{dispatch:dispatch}

	return handler
}

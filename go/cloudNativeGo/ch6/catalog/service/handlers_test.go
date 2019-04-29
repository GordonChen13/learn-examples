package service

import (
	"encoding/json"
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

func TestHandler_Get(t *testing.T) {
	engine = getEngine()

	t.Run("get catalog by sku", func(t *testing.T) {
		sku := 	"THINGAMAJIG12"

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/catalog?sku=" + sku, nil)
		engine.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)

		got := &Catalog{}
		json.NewDecoder(res.Body).Decode(got)

		assert.Equal(t, sku, got.SKU)
		assert.Equal(t, 100, got.QuantityInStock)
		assert.Equal(t, 14, got.ShipsWithin)
	})

}

func TestHandler_Create(t *testing.T) {
	engine = getEngine()

	t.Run("return 400 when post empty form", func(t *testing.T) {
		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/catalog", nil)
		engine.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("return 400 when post form without stock", func(t *testing.T) {
		form := url.Values{}
		form.Set("sku", "XiaoMi 9")
		form.Set("shipswithin", "20")
		body := strings.NewReader(form.Encode())

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/catalog", body)
		engine.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("return 400 when post form without sku", func(t *testing.T) {
		form := url.Values{}
		form.Set("shipswithin", "20")
		form.Set("stock", "5")
		body := strings.NewReader(form.Encode())

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/catalog", body)
		engine.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("return 400 when post form without shipswithin", func(t *testing.T) {
		form := url.Values{}
		form.Set("sku", "XiaoMi 9")
		form.Set("stock", "5")
		body := strings.NewReader(form.Encode())

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/catalog", body)
		engine.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("return catalog when created", func(t *testing.T) {
		form := url.Values{}
		form.Set("sku", "XiaoMi 9")
		form.Set("shipswithin", "20")
		form.Set("stock", "5")
		body := strings.NewReader(form.Encode())

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/catalog", body)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		engine.ServeHTTP(res, req)

		catalog := &Catalog{}
		assert.Equal(t, http.StatusCreated, res.Code)
		json.NewDecoder(res.Body).Decode(catalog)

		assert.Equal(t, "XiaoMi 9", catalog.SKU)
		assert.Equal(t, "20", strconv.Itoa(catalog.ShipsWithin))
		assert.Equal(t, "5", strconv.Itoa(catalog.QuantityInStock))
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
	NewHandler(engine)

	return engine
}


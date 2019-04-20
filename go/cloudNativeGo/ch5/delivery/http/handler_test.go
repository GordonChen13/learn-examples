package http

import (
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch5/usecase/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var mockServer *gin.Engine = nil

func TestTest(t *testing.T) {
	res := serveRequest(t, http.MethodGet, "/test", nil)



	got := res.Body.String()
	want := "This is a test"

	assert.Equal(t, res.Code, 200)
	assert.Equal(t, got, want)
}

func TestCreateMatch(t *testing.T) {
	name := "gordon"
	data := url.Values{}
	data.Set("name", name)

	res := serveRequest(t, http.MethodPost, "/match", strings.NewReader(data.Encode()))
	got := res.Body.String()
	want := name

	assert.Equal(t, res.Code, 201)
	assert.Equal(t, got, want)

}

func serveRequest(t *testing.T, method, url string, body io.Reader) *httptest.ResponseRecorder {
	server := getMockServer(t)

	req, _ := http.NewRequest(method, url, body)
	res := httptest.NewRecorder()

	switch method {
	case http.MethodPost:
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	server.ServeHTTP(res, req)
	return  res
}

func newMockServer(t *testing.T) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	ctrl := gomock.NewController(t)
	mockMatchUseCase := mock_usecase.NewMockMatch(ctrl)
	NewMatchHandler(router, mockMatchUseCase)

	return router
}

func getMockServer(t *testing.T) *gin.Engine {
	if mockServer != nil {
		return mockServer
	}
	mockServer = newMockServer(t)
	return mockServer
}

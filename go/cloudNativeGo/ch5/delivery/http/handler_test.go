package http

import (
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch5/models"
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch5/usecase/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var router *gin.Engine = nil

func TestTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	_, res, engine, _ := initHandler(t)

	req := makeRequest(t, http.MethodGet, "/test", nil)

	engine.ServeHTTP(res, req)

	got := res.Body.String()
	want := "This is a test"

	assert.Equal(t, res.Code, 200)
	assert.Equal(t, got, want)
}

func TestCreateMatch(t *testing.T) {
	name := "gordon"

	ctx, res, engine, mockUseCase := initHandler(t)

	id := models.NewMatchId()
	match := &models.Match{
		Id:    id,
		Name:  name,
		Moves: nil,
	}
	mockUseCase.EXPECT().Create(ctx, match)

	data := url.Values{}
	data.Set("name", name)
	req := makeRequest(t, http.MethodPost, "/match", strings.NewReader(data.Encode()))

	engine.ServeHTTP(res, req)

	got := res.Body.String()
	want := name

	assert.Equal(t, res.Code, 201)
	assert.Equal(t, got, want)

}

func makeRequest(t *testing.T, method, url string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, url, body)

	switch method {
	case http.MethodPost:
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	return req
}

func initHandler(t *testing.T) (ctx context.Context, res *httptest.ResponseRecorder, engine *gin.Engine, mockUseCase *mock_usecase.MockMatch) {
	ctrl := gomock.NewController(t)
	gin.SetMode(gin.ReleaseMode)
	defer ctrl.Finish()

	res = httptest.NewRecorder()
	ctx, engine = gin.CreateTestContext(res)
	mockUseCase = mock_usecase.NewMockMatch(ctrl)
	NewMatchHandler(engine, mockUseCase)

	return
}

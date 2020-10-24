package v1

import (
	"github.com/GordonChen13/learn-examples/go/blog-service/global"
	"github.com/GordonChen13/learn-examples/go/blog-service/pkg/app"
	"github.com/GordonChen13/learn-examples/go/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (t Article) Get(c *gin.Context) {
	global.Logger.Error("get article", zap.String("name", "chen"), zap.Int("age", 18))
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

func (t Article) List(c *gin.Context) {

}

func (t Article) Create(c *gin.Context) {

}

func (t Article) Update(c *gin.Context) {

}

func (t Article) Delete(c *gin.Context) {

}

package v1

import (
	"github.com/GordonChen13/learn-examples/go/blog-service/pkg/app"
	"github.com/GordonChen13/learn-examples/go/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (t Article) Get(c *gin.Context) {
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

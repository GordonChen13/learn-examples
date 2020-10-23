package routers

import (
	v1 "github.com/GordonChen13/learn-examples/go/blog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	tag := v1.NewTag()
	article := v1.NewArticle()
	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/tags", tag.Create)
		apiV1.GET("/tags/:id", tag.Get)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tgtags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.Update)
		apiV1.GET("/tags", tag.List)

		apiV1.POST("/articles", article.Create)
		apiV1.GET("/articles/:id", article.Get)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.PATCH("/articles/:id/state", article.Update)
		apiV1.GET("/articles", article.List)
	}

	return r
}

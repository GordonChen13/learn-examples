package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Controller *Controller
}
// register route and init handler
func NewHandler(router *gin.Engine){
	controller := NewController()
	handler := &Handler{
		Controller:controller,
	}

	router.GET("/catalog", handler.Get)
	router.POST("/catalog", handler.Create)
}

func (h *Handler) Get(ctx *gin.Context) {
	res, err := h.Controller.Get(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) Create(ctx *gin.Context) {
	var form Catalog
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Controller.Create(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, res)
}

package v1

import (
	"github.com/GordonChen13/learn-examples/go/blog-service/global"
	"github.com/GordonChen13/learn-examples/go/blog-service/internal/service"
	"github.com/GordonChen13/learn-examples/go/blog-service/pkg/app"
	"github.com/GordonChen13/learn-examples/go/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0,1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	res := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Error("Tag List", zap.Strings("errs", errs.Errors()))
		res.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	total, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Error("svc.CountTag", zap.String("err", err.Error()))
		res.ToResponse(errcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Error("svc.GetTagList", zap.String("err", err.Error()))
		res.ToResponse(errcode.ErrorGetTagListFail)
		return
	}

	res.ToResponseList(tags, total)
	return
}

// @Summary 新增标签
// @Produce json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0,1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	res := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Error("Tag Create app.BindAndValid(c, &param)", zap.Any("param", param), zap.Strings("errs", errs.Errors()))
		res.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Error("Tag Create svc.CreateTag(&param)", zap.Any("param", param), zap.Strings("errs", errs.Errors()))
		res.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	res.ToResponse(gin.H{})
	return
}

// @Summary 更新标签
// @Produce json
// @Param id path int true "标签ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0,1) default(1)
// @Param modified_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {

}

// @Summary 删除标签
// @Produce  json
// @Param id path int true "标签ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {

}

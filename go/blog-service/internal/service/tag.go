package service

import (
	"github.com/GordonChen13/learn-examples/go/blog-service/internal/model"
	"github.com/GordonChen13/learn-examples/go/blog-service/pkg/app"
)

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=32"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=32"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=32"`
	CreatedBy string `form:"created_by" binding:"min=3,max=32"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID        uint   `form:"id" binding:"required,gte=1"`
	Name      string `form:"name" binding:"required,min=3,max=32"`
	UpdatedBy string `form:"updated_by" binding:"required,min=3,max=32"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteTagRequest struct {
	ID uint `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountTag(param *CountTagRequest) (int64, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.UpdatedBy)
}

func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}

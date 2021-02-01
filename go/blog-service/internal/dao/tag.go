package dao

import (
	"github.com/GordonChen13/learn-examples/go/blog-service/internal/model"
	"github.com/GordonChen13/learn-examples/go/blog-service/pkg/app"
)

func (d *Dao) CountTag(name string, state uint8) (int64, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint, name string, state uint8, updatedBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{
			UpdatedBy: updatedBy,
		},
	}
	tag.ID = id
	return tag.Update(d.engine, tag)
}

func (d *Dao) DeleteTag(id uint) error {
	tag := model.Tag{}
	tag.ID = id
	return tag.Delete(d.engine)
}

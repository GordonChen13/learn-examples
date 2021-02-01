package model

import (
	"github.com/GordonChen13/learn-examples/go/blog-service/pkg/app"
	"gorm.io/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name" gorm:"type:VARCHAR(32) DEFAULT '' COMMENT '标签名称'"`
	State uint8  `json:"state" gorm:"type:TINYINT(4) UNSIGNED DEFAULT '1' COMMENT '状态 0:禁用 1:启I用'"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffSet, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffSet > 0 && pageSize > 0 {
		db = db.Offset(pageOffSet).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Model(&t).Find(&tags).Error; err != nil {
		return tags, err
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ?", t.ID).Updates(values).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ?", t.ID).Delete(&t).Error
}

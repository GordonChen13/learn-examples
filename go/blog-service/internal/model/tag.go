package model

import "github.com/GordonChen13/learn-examples/go/blog-service/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name" gorm:"type:VARCHAR(32) DEFAULT '' COMMENT '标签名称'"`
	State uint8  `json:"state" gorm:"type:TINYINT(4) UNSIGNED DEFAULT '1' COMMENT '状态 0:禁用 1:启I用'"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (a Tag) TableName() string {
	return "blog_tag"
}

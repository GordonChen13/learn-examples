package model

import "github.com/GordonChen13/learn-examples/go/blog-service/pkg/app"

type Article struct {
	*Model
	Title         string `json:"title" gorm:"type:VARCHAR(128) DEFAULT '' COMMENT '标题'"`
	Desc          string `json:"desc" gorm:"type:VARCHAR(255) DEFAULT '' COMMENT '描述'"`
	Content       string `json:"content" gorm:"type:TEXT COMMENT '内容'"`
	CoverImageUrl string `json:"cover_image_url" gorm:"type:VARCHAR(128) DEFAULT '' COMMENT '封面链接'"`
	State         uint8  `json:"state" gorm:"type:TINYINT(4) UNSIGNED DEFAULT '1' COMMENT '状态 0:禁用 1:启I用'"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager app.Pager
}

func (a *Article) TableName() string {
	return "blog_article"
}

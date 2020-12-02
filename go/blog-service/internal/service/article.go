package service

type CountArticleRequest struct {
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	Name  string `form:"name" binding:"max=32"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Name  string `form:"name" binding:"required,min=3,max=32"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	Name  string `form:"name" binding:"required,min=3,max=32"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

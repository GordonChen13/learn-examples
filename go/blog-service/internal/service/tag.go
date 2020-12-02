package service

type CountTagRequest struct {
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Title string `form:"title" binding:"max=128"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Title         string `form:"title" binding:"required,min=3,max=128"`
	Desc          string `form:"desc" binding:"required,min=3,max=256"`
	Content       string `form:"content" binding:"required,min=3"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,min=3,max=128"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"required,min=3,max=128"`
	Desc          string `form:"desc" binding:"required,min=3,max=256"`
	Content       string `form:"content" binding:"required,min=3"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,min=3,max=128"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

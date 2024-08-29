package service

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=01"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=01"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"requirded,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"requirded,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=01"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"requirded,gte=1"`
	Name       string `form:"name" binding:"min=3,max=100"`
	State      uint8  `form:"state" binding:"requirded,oneof=01"`
	ModifiedBy string `form:"modified_by" binding:"requirded,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"requirded,gte=1"`
}

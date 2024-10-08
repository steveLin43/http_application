package service

import (
	"http_application/internal/model"
	"http_application/pkg/app"
)

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

func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}

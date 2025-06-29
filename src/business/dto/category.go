package dto

import "github.com/NupalHariz/DD/src/business/entity"

type CreateCategoryParam struct {
	Name string `json:"name" binding:"required"`
}

func (c *CreateCategoryParam) ToCategoryInputParam(userId int64) entity.CategoryInputParam {
	return entity.CategoryInputParam{
		Name:   c.Name,
		UserId: userId,
	}
}

type GetAllCategoryResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

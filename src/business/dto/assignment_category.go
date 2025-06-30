package dto

import "github.com/NupalHariz/DD/src/business/entity"

type CreateAssignmentCategory struct {
	Name string `json:"name"`
}

func (c *CreateAssignmentCategory) ToAssignmentCategoryInputParam(userId int64) entity.AssignmentCategoryInputParam {
	return entity.AssignmentCategoryInputParam{
		UserId: userId,
		Name:   c.Name,
	}
}

type GetAllAssignmentCategoryResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

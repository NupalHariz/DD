package dto

import (
	"github.com/NupalHariz/DD/src/business/entity"
)

type CreateAssignmentParam struct {
	CategoryId int64  `json:"category_id"`
	Name       string `json:"name"`
	Deadline   string `json:"deadline"`
	Status     string `json:"status"`
	Priority   string `json:"priority"`
}

func (c *CreateAssignmentParam) ToAssignmentInputParam(userId int64) entity.AssignmentInputParam {
	return entity.AssignmentInputParam{
		UserId:     userId,
		CategoryId: c.CategoryId,
		Name:       c.Name,
		Deadline:   c.Deadline,
		Status:     entity.Status(c.Status),
		Priority:   entity.Priority(c.Priority),
	}
}

type UpdateAssignmentParam struct {
	Id         int64  `json:"-" uri:"id"`
	CategoryId int64  `json:"category_id"`
	Name       string `json:"name"`
	Deadline   string `json:"deadline"`
	Status     string `json:"status"`
	Priority   string `json:"priority"`
}

func (c *UpdateAssignmentParam) ToAssignmentUpdateParam() entity.AssignmentUpdateParam {
	return entity.AssignmentUpdateParam{
		CategoryId: c.CategoryId,
		Name:       c.Name,
		Deadline:   c.Deadline,
		Status:     entity.Status(c.Status),
		Priority:   entity.Priority(c.Priority),
	}
}

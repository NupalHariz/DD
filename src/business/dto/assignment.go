package dto

import (
	"time"

	"github.com/NupalHariz/DD/src/business/entity"
)

type CreateAssignmentParam struct {
	CategoryId int64  `json:"category_id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Deadline   string `json:"deadline" binding:"required"`
	Status     string `json:"status" binding:"required"`
	Priority   string `json:"priority" binding:"required"`
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

type GetAllAssignmentParam struct {
	CategoryId int64  `query:"category_id"`
	Status     string `query:"status"`
	Priority   string `query:"priority"`
	entity.PaginationParam
}

type GetAllAssignmentResponse struct {
	Id       int64     `json:"id"`
	Category string    `json:"category"`
	Name     string    `json:"name"`
	Deadline time.Time `json:"deadline"`
	Status   string    `json:"status"`
	Priority string    `json:"priority"`
}

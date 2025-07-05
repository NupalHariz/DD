package dto

import (
	"github.com/NupalHariz/DD/src/business/entity"
	"github.com/reyhanmichiels/go-pkg/v2/null"
)

type CreateDailyAssignmentParam struct {
	Name string `json:"name" binding:"required"`
}

func (c *CreateDailyAssignmentParam) ToDailyAssignmentInputParam(userId int64) entity.DailyAssignmentInputParam {
	return entity.DailyAssignmentInputParam{
		UserId: userId,
		Name:   c.Name,
	}
}

type UpdateDailyAssignmentParam struct {
	Id     int64  `json:"-" uri:"id"`
	Name   string `json:"name"`
	IsDone bool   `json:"is_done"`
}

func (u *UpdateDailyAssignmentParam) ToDailyAssignmentUpdateParam() entity.DailyAssignmentUpdateParam {
	return entity.DailyAssignmentUpdateParam{
		Name:   u.Name,
		IsDone: null.BoolFrom(u.IsDone),
	}
}

type GetAllDailyAssignmentResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	IsDone bool   `json:"is_done"`
}

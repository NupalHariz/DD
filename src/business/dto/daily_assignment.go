package dto

import "github.com/NupalHariz/DD/src/business/entity"

type CreateDailyAssignmentParam struct {
	Name string `json:"name"`
}

func (c *CreateDailyAssignmentParam) ToDailyAssignmentInputParam(userId int64) entity.DailyAssignmentInputParam {
	return entity.DailyAssignmentInputParam{
		UserId: userId,
		Name: c.Name,
	}
}

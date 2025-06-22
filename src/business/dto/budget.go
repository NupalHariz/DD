package dto

import "github.com/NupalHariz/DD/src/business/entity"

type CreateBudgetParam struct {
	CategoryId int64  `json:"category_id"`
	Amount     int64  `json:"amount"`
	Type       string `json:"type"`
}

func (c *CreateBudgetParam) ToBudgetInputParam(userId int64) entity.BudgetInputParam {
	return entity.BudgetInputParam{
		UserId:     userId,
		CategoryId: c.CategoryId,
		Amount:     c.Amount,
		Type:       entity.BudgetType(c.Type),
	}
}

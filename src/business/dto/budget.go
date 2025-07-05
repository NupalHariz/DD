package dto

import "github.com/NupalHariz/DD/src/business/entity"

type CreateBudgetParam struct {
	CategoryId int64  `json:"category_id" binding:"required"`
	Amount     int64  `json:"amount" binding:"required, gte= 0"`
	Type       string `json:"type" binding:"required"`
}

func (c *CreateBudgetParam) ToBudgetInputParam(userId int64) entity.BudgetInputParam {
	return entity.BudgetInputParam{
		UserId:     userId,
		CategoryId: c.CategoryId,
		Amount:     c.Amount,
		Type:       entity.BudgetType(c.Type),
	}
}

type UpdateBudgetParam struct {
	Id     int64  `json:"-" uri:"id"`
	Amount int64  `json:"amount" binding:"gte=0"`
	Type   string `json:"type"`
}

func (u *UpdateBudgetParam) ToBudgetUpdateParam() entity.BudgetUpdateParam {
	return entity.BudgetUpdateParam{
		Amount: u.Amount,
		Type:   entity.BudgetType(u.Type),
	}
}

type GetAllBudgetResponse struct {
	Id             int64             `db:"id"`
	Category       string            `db:"category"`
	Amount         int64             `db:"amount"`
	CurrentExpense int64             `db:"current_expense"`
	Type           entity.BudgetType `db:"time_period"`
}

type GetBudgetParam struct {
	Type string `json:"type"`
	entity.PaginationParam
}

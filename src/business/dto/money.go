package dto

import "github.com/NupalHariz/DD/src/business/entity"

type CreateTransactionParam struct {
	Amount     int64  `json:"amount" binding:"required"`
	CategoryId int64  `json:"category_id" binding:"required"`
	Type       string `json:"type" binding:"required"`
}

func (c *CreateTransactionParam) ToInputMoneyParam(userId int64) entity.MoneyInputParam {
	return entity.MoneyInputParam{
		UserId:     userId,
		Amount:     c.Amount,
		CategoryId: c.CategoryId,
		Type:       entity.MoneyType(c.Type),
	}
}

type UpdateTransactionParam struct {
	Id         int64  `json:"-" uri:"id"`
	Amount     int64  `json:"amount"`
	CategoryId int64  `json:"category_id"`
	Type       string `json:"type"`
}

func (u *UpdateTransactionParam) ToMoneyUpdateParam() entity.MoneyUpdateParam {
	return entity.MoneyUpdateParam{
		Amount:     u.Amount,
		CategoryId: u.CategoryId,
		Type:       entity.MoneyType(u.Type),
	}
}

type GetTransactionParam struct {
	CategoryId int64  `form:"category_id"`
	Type       string `form:"type"`
	entity.PaginationParam
}

type GetTransactionResponse struct {
	Id       int64            `json:"id"`
	Amount   int64            `json:"amount"`
	Category string           `json:"category"`
	Type     entity.MoneyType `json:"type"`
}

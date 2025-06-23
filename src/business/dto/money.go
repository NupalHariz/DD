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

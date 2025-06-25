package entity

import "github.com/reyhanmichiels/go-pkg/v2/query"

type BudgetType string

const (
	Weekly  BudgetType = "WEEKLY"
	Monthly BudgetType = "MONTHLY"
)

type Budget struct {
	Id             string     `db:"id"`
	UserId         int64      `db:"user_id"`
	CategoryId     int64      `db:"category_id"`
	Amount         int64      `db:"amount"`
	CurrentExpense int64      `db:"current_expense"`
	Type           BudgetType `db:"time_period"`
}

type BudgetInputParam struct {
	UserId     int64      `db:"user_id"`
	CategoryId int64      `db:"category_id"`
	Amount     int64      `db:"amount"`
	Type       BudgetType `db:"type"`
}

type BudgetUpdateParam struct {
	UserId         int64      `db:"user_id"`
	CategoryId     int64      `db:"category_id"`
	CurrentExpense int64      `db:"current_expense"`
	Amount         int64      `db:"amount"`
	Type           BudgetType `db:"time_period"`
}

type BudgetUpdateParamm struct {
	Amount         int64      `db:"amount"`
	Type           BudgetType `db:"time_period"`
}

type BudgetParam struct {
	Id     int64 `db:"id" param:"id"`
	Option query.Option
}

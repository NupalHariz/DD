package entity

import "github.com/reyhanmichiels/go-pkg/v2/query"

type MoneyType string

const (
	Income  MoneyType = "Income"
	Expense MoneyType = "Expense"
)

type Money struct {
	Id         int64     `db:"id"`
	UserId     int64     `db:"user_id"`
	Amount     int64     `db:"amount"`
	CategoryId int64     `db:"category_id"`
	Type       MoneyType `db:"type"`
}

type MoneyInputParam struct {
	UserId     int64     `db:"user_id"`
	Amount     int64     `db:"amount"`
	CategoryId int64     `db:"category_id"`
	Type       MoneyType `db:"type"`
}

type MoneyUpdateParam struct {
	Amount     int64     `db:"amount"`
	CategoryId int64     `db:"category_id"`
	Type       MoneyType `db:"type"`
}

type MoneyParam struct {
	Id     int64 `db:"id" param:"id"`
	Option query.Option
}

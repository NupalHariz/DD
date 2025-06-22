package entity

type MoneyType string

const (
	Income  MoneyType = "Income"
	Expense MoneyType = "Expense"
)

type Money struct {
	Id         int64     `db:"id"`
	UserId     int64    `db:"user_id"`
	Amount     int64     `db:"amount"`
	CategoryId int64     `db:"category_id"`
	Type       MoneyType `db:"type"`
}

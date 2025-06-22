package entity

type BudgetType string

const (
	Weekly  BudgetType = "WEEKLY"
	Monthly BudgetType = "MONTHLY"
)

type Budget struct {
	Id         string     `db:"id"`
	UserId     int64      `db:"user_id"`
	CategoryId int64      `db:"category_id"`
	Amount     int64      `db:"amount"`
	Type       BudgetType `db:"type"`
}

type BudgetInputParam struct {
	UserId     int64      `db:"user_id"`
	CategoryId int64      `db:"category_id"`
	Amount     int64      `db:"amount"`
	Type       BudgetType `db:"type"`
}

package entity

import "time"

type HistoryBudget struct {
	Id          int64      `db:"id"`
	UserId      int64      `db:"user_id"`
	BudgetId    int64      `db:"budget_id"`
	CategoryId  int64      `db:"category_id"`
	Spent       int64      `db:"spent"`
	Planned     int64      `db:"planned"`
	Type        BudgetType `db:"type"`
	PeriodStart time.Time  `db:"period_start"`
	PeriodEnd   time.Time  `db:"period_end"`
	CreatedAt   time.Time  `db:"created_at"`
}

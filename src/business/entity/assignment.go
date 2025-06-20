package entity

import "time"

type Status string
type Priority string

const (
	OnGoing Status = "ONGOING"
	Done    Status = "DONE"

	Low    Priority = "LOW"
	Medium Priority = "MEDIUM"
	High   Priority = "HIGH"
)

type Assignment struct {
	Id         int64     `db:"id"`
	UserId     string    `db:"user_id"`
	Name       string    `db:"name"`
	Deadline   time.Time `db:"deadline"`
	Status     Status    `db:"status"`
	Priority   Priority  `db:"priority"`
	CategoryId int64     `db:"category_id"`
}

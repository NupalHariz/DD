package entity

import (
	"time"

	"github.com/reyhanmichiels/go-pkg/v2/query"
)

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
	UserId     int64     `db:"user_id"`
	Name       string    `db:"name"`
	Deadline   time.Time `db:"deadline"`
	Status     Status    `db:"status"`
	Priority   Priority  `db:"priority"`
	CategoryId int64     `db:"category_id"`
}

type AssignmentInputParam struct {
	UserId     int64    `db:"user_id"`
	CategoryId int64    `db:"category_id"`
	Name       string   `db:"name"`
	Deadline   string   `db:"deadline"`
	Status     Status   `db:"status"`
	Priority   Priority `db:"priority"`
}

type AssignmentUpdateParam struct {
	CategoryId int64    `db:"category_id"`
	Name       string   `db:"name"`
	Deadline   string   `db:"deadline"`
	Status     Status   `db:"status"`
	Priority   Priority `db:"priority"`
}

type AssignmentParam struct {
	Id          int64  `db:"id" param:"id"`
	UserId      int64  `db:"user_id" param:"user_id"`
	Deadline    string `db:"deadline" param:"deadline"`
	Status      string `db:"status" param:"status"`
	Option      query.Option
	BypassCache bool
	PaginationParam
}

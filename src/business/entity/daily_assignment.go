package entity

import (
	"github.com/reyhanmichiels/go-pkg/v2/null"
	"github.com/reyhanmichiels/go-pkg/v2/query"
)

type DailyAssignment struct {
	Id     int64  `db:"id"`
	UserId int64  `db:"user_id"`
	Name   string `db:"name"`
	IsDone bool   `db:"is_done"`
}

type DailyAssignmentInputParam struct {
	UserId int64  `db:"user_id"`
	Name   string `db:"name"`
}

type DailyAssignmentUpdateParam struct {
	Name   string    `db:"name"`
	IsDone null.Bool `db:"is_done"`
}

type DailyAssignmentParam struct {
	Id          int64 `db:"id" param:"id"`
	UserId      int64 `db:"user_id" param:"user_id"`
	Option      query.Option
	BypassCache bool
}

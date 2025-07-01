package entity

import "github.com/reyhanmichiels/go-pkg/v2/query"

type AssignmentCategory struct {
	Id     int64  `db:"id"`
	UserId int64  `db:"user_id"`
	Name   string `db:"name"`
}

type AssignmentCategoryInputParam struct {
	UserId int64  `db:"user_id"`
	Name   string `db:"name"`
}

type AssignmentCategoryParam struct {
	Ids         []int64 `db:"id" param:"id"`
	UserId      int64   `db:"user_id" param:"user_id"`
	Option      query.Option
	BypassCache bool
}

package entity

import "github.com/reyhanmichiels/go-pkg/v2/query"

type Category struct {
	Id     int64  `db:"id"`
	UserId int64  `db:"user_id"`
	Name   string `db:"name"`
}

type CategoryInputParam struct {
	UserId int64  `db:"user_id"`
	Name   string `db:"name"`
}

type CategoryParam struct {
	Ids    []int64 `db:"id" param:"id"`
	UserId int64   `db:"user_id" param:"user_id"`
	Option query.Option
}

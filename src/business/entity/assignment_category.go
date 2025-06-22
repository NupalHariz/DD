package entity

type AssignmentCategory struct {
	Id     int64  `db:"id"`
	UserId int64 `db:"user_id"`
	Name   string `db:"name"`
}

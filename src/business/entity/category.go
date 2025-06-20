package entity

type Category struct {
	Id     int64  `db:"id"`
	UserId string `db:"user_id"`
	Name   string `db:"name"`
}

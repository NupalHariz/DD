package entity

type DailyAssignment struct {
	Id     int64  `db:"id"`
	UserId string `db:"user_id"`
	Name   string `db:"name"`
	IsDone bool   `db:"is_done"`
}

package entity

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

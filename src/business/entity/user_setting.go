package entity

import "time"

type UserSetting struct {
	UserId        string    `db:"user_id"`
	DailyReminder time.Time `db:"daily_reminder"`
	ReminderHour  time.Time `db:"reminder_hour"`
}

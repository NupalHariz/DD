package entity

import "time"

type UserSetting struct {
	UserId        int64    `db:"user_id"`
	DailyReminder time.Time `db:"daily_reminder"`
	ReminderHour  time.Time `db:"reminder_hour"`
}

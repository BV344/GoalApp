package domain

import "time"

type ScheduleSlot struct {
	ID        int64     `json:"id"`
	GoalID    int64     `json:"goal_id"`
	Day       string    `json:"day"`        // monday, tuesday, ...
	StartTime string    `json:"start_time"` // "09:00"
	EndTime   string    `json:"end_time"`   // "10:00"
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}

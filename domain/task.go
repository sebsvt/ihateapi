package domain

import "time"

type Task struct {
	TaskID    string    `bson:"task_id"`
	UserID    string    `bson:"user_id"`
	CreatedAt time.Time `bson:"created_at"`
}

package domain

import "time"

type UpdateTaskRequest struct {
	Title       string
	Description string
	Deadline    time.Time
}

package domain

import "time"

type FetchTaskRequest struct {
	UserID uint
	Offset int
	Limit  int
}

type FetchTaskResponse struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      uint      `json:"user_id"`
}

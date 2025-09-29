package response

import "time"

type TaskResponse struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskListResponse struct {
	Tasks []*TaskResponse
}

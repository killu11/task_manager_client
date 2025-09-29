package request

type CreateTaskRequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	StatusID    int    `json:"status_id,omitempty"`
}

func NewCreateTaskRequest(title, description string, statusID int) *CreateTaskRequest {
	return &CreateTaskRequest{
		Title:       title,
		Description: description,
		StatusID:    statusID,
	}
}

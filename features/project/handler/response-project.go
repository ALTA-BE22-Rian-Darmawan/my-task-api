package handler

type ProjectResponse struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	ProjectName string `json:"project_name"`
	Description string `json:"Description"`
}

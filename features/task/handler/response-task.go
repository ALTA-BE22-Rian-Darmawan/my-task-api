package handler

type TaskResponse struct {
	ID          uint   `json:"id"`
	ProjectID   uint   `json:"project_id"`
	TaskName    string `json:"project_name"`
	Description string `json:"Description"`
}

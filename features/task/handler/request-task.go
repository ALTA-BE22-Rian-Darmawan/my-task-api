package handler

type TaskRequest struct {
	IDProject   uint   `json:"project_id" form:"project_id"`
	TaskName    string `json:"task_name" form:"task_name"`
	Description string `json:"description" form:"description"`
	Status      string `json:"status" form:"status"`
}

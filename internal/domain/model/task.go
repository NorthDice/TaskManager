package model

// TaskList represents a task in the task management system.
type TaskList struct {
	Id          int    `json:"id" bson:"id"`
	UserId      string `json:"user_id" bson:"user_id"`
	Title       string `json:"title" binding:"required" bson:"title"`
	Description string `json:"description" bson:"description"`
}

// UpdateTaskListInput is used to update a task list's title and description.
type UpdateTaskListInput struct {
	Title       *string `json:"title" bson:"title"`
	Description *string `json:"description" bson:"description"`
}

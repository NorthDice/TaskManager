package repository

type TaskListMongo struct {
	// MongoDB client or session can be added here
}

func NewTaskListMongo() *TaskListMongo {
	return &TaskListMongo{
		// Initialize MongoDB client or session here if needed
	}
}

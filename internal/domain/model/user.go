package model

// User represents a user in the task management system.
type User struct {
	Id       int    `json:"-" bson:"id"`
	Username string `json:"username" binding:"required" bson:"username"`
	Password string `json:"password" binding:"required" bson:"password"`
	Email    string `json:"email" binding:"required,email" bson:"email"`
}

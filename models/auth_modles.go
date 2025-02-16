package models

type User struct {
	Name     string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password"`
}

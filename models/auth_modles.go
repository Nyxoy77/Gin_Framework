package models

type User struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password"`
}

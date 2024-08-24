package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"user_name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

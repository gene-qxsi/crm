package models

type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

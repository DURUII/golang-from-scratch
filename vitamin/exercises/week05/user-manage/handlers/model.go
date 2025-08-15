package handlers

type User struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" binding:"required" form:"email"`
	Age   int    `json:"age" form:"age"`
}

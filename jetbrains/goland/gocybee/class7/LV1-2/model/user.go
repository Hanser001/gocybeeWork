package model

type User struct {
	Id       int
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Question string
	Answer   string
}

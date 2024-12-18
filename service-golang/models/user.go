package models

import "gorm.io/gorm"

type User struct{
	gorm.Model

	Username string `form:"username" json:"username"`
	Email string `form:"email" json:"email"`
	Name string `form:"name" json:"name"`
	Age string `form:"age" json:"age"`
	Address string `form:"address" json:"address"`
}
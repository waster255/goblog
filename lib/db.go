package lib

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

type Passages struct {
	Title string
	User User `gorm:"constraint:OnDelete:CASCADE"`
	UserId int
	Content string
}
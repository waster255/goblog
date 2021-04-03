package lib

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func AddUser(c *gin.Context) {

	var user User

	if c.ShouldBind(&user) == nil {
		log.Println(user.Username)
		log.Println(user.Password)
	}

	c.JSON(200, gin.H{
		"status": "added",
		"username": user.Username,
	})
	
	session, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to conenct database")
	}
	session.AutoMigrate(&User{})
	session.Create(&User{Username: user.Username, Password: user.Password})
}
package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name         string `json:"name"`
	Login        string `json:"login"`
	Password 	 string `json:"password"`
	CreatedAt    time.Time `json:"createdAt"`
}

func main() {

	user := User{
		Name:         "Maciek",
		Login:        "mgaszynski",
		Password: "$2a$10$eRBBlboWvbDd4YX.wlFcI.jnwo7iN.40XzM4rWgieohwL1fRQZZI.",
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"user":    user,
		})
	})
	r.POST("/login", func(c *gin.Context) {
		newUser := User{}
		c.Bind(&newUser) 
		
		var response string

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(newUser.Password)); err != nil {
			response = "Wrong credentials"
		} else {
			response = "Logged in"
		}

		c.String(200, response)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

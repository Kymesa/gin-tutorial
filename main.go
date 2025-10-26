package main

import (
	"net/http"

	"gin-tutorial/config"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	app := gin.Default()

	app.GET("/", func(response *gin.Context) {
		response.JSON(http.StatusOK, gin.H{
			"message": "Bienvenido a mi API con Gin 🚀",
		})
	})

	app.GET("/users", func(response *gin.Context) {
		users := []User{
			{Id: 1, Name: "keiner mesa",
				Age: 21},
			{Id: 2, Name: "keiner yesid",
				Age: 21},
		}
		response.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	})

	app.Run(config.PORT)
}

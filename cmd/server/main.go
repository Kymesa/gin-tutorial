package main

import (
	"gin-tutorial/internal/auth"
	"gin-tutorial/internal/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {

		if os.Getenv("DEV") == "TRUE" {

			log.Fatal("Error loading .env file")
		}

	}

	database.ConnectDB()

	database.DB.AutoMigrate(&auth.User{})

	app := gin.Default()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	v1 := app.Group("api/v1")
	{
		v1.POST("/register", auth.Register)
		v1.POST("/login", auth.Login)

	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Escuchando en el puerto %s", port)
	app.Run(":" + port)
}

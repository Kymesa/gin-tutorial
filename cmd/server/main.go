package main

import (
	"gin-tutorial/internal/auth"
	"gin-tutorial/internal/database"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://valleinmuebles.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	v1 := app.Group("api/v1")
	{
		v1.GET("/", auth.Test)
		v1.POST("/login", auth.Login)
		v1.POST("/register", auth.Register)

	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Escuchando en el puerto %s", port)
	app.Run(":" + port)
}

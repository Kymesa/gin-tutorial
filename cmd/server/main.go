package main

import (
	// "gin-tutorial/internal/auth"
	"gin-tutorial/internal/book"
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
	//  Conectamos a la base de datos PostgreSQL
	database.ConnectDB()

	// Migramos el modelo Book (si no existe la tabla, la crea automáticamente)
	database.DB.AutoMigrate(&book.Book{})
	// database.DB.AutoMigrate(&auth.User{})

	// Creamos una nueva instancia del servidor Gin
	app := gin.Default()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	v1 := app.Group("api/v1")
	// {
	// 	v1.POST("/register", auth.Register)
	// 	v1.POST("/login", auth.Login)
	// 	v1.POST("/refresh", auth.Refresh)
	// }

	// BOOKS
	repo := book.NewRepository()
	service := book.NewService(repo)
	handler := book.NewHandler(service)
	book.SetupRouter(handler, v1)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Escuchando en el puerto %s", port)
	app.Run(":" + port)
}

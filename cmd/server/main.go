package main

import (
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

	// Creamos una nueva instancia del servidor Gin
	app := gin.Default()

	// Registramos las rutas del módulo "book"
	book.RegisterRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Escuchando en el puerto %s", port)
	app.Run(":" + port)
}

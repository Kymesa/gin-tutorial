package main

import (
	"gin-tutorial/config/jwt"
	"gin-tutorial/internal/book"
	"gin-tutorial/internal/database"
	"log"
	"net/http"
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

	v1 := app.Group("/api/v1")

	// Ruta pública: login
	app.POST("/api/v1/login", func(c *gin.Context) {
		var req struct {
			UserID   string `json:"userId"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}

		if req.Password != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
			return
		}

		token, _ := jwt.GenerateJWT(req.UserID)
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	// BOOKS
	book.RegisterRoutes(v1)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Escuchando en el puerto %s", port)
	app.Run(":" + port)
}

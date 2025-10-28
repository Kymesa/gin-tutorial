package main

import (
	"gin-tutorial/internal/book"

	"gin-tutorial/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conectamos a la base de datos PostgreSQL
	database.ConnectDB()

	// Migramos el modelo Book (si no existe la tabla, la crea automáticamente)
	database.DB.AutoMigrate(&book.Book{})

	// Creamos una nueva instancia del servidor Gin
	app := gin.Default()

	// Registramos las rutas del módulo "book"
	book.RegisterRoutes(app)

	// Iniciamos el servidor en el puerto 8080
	app.Run(":8080")
}

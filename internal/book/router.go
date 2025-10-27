package book

import "github.com/gin-gonic/gin"

// RegisterRoutes registra las rutas del módulo de libros
func RegisterRoutes(r *gin.Engine) {
	repo := NewRepository()
	handler := NewHandler(repo)

	books := r.Group("/books")
	{
		books.POST("/", handler.CreateBook)
		books.GET("/", handler.GetBooks)
		books.GET("/:id", handler.GetBookByID)
		books.PUT("/:id", handler.UpdateBook)
		books.DELETE("/:id", handler.DeleteBook)
	}
}

package book

import (
	"gin-tutorial/config/jwt"
	
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	repo := NewRepository()
	services := NewService(repo)

	books := router.Group("/books")
	books.Use(jwt.AuthMiddleware())
	{
		books.POST("/", services.CreateBook)
		books.GET("/", services.GetBooks)
		books.GET("/:id", services.GetBookByID)
		books.PUT("/:id", services.UpdateBook)
		books.DELETE("/:id", services.DeleteBook)
	}
}

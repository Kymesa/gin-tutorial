package book

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *BookHandler, v1 *gin.RouterGroup) {

	{
		books := v1.Group("/books")
		{
			// books.POST("/", services.CreateBook)
			// books.GET("/", services.GetBooks)
			// books.GET("/:id", services.GetBookByID)
			// books.PUT("/:id", services.UpdateBook)
			books.DELETE("/:id", userHandler.DeleteBookHandler)
		}
	}
}

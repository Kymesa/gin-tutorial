package book

import (
	"gin-tutorial/config/res"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service IService
}

func NewHandler(service IService) *BookHandler {
	return &BookHandler{service}
}

func (h *BookHandler) DeleteBookHandler(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		res.Error(context, http.StatusInternalServerError, "No se pudo obtenido el id", nil)
		return
	}

	if err := h.service.DeleteBookService(uint(id)); err != nil {

		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Libro eliminado"})
}

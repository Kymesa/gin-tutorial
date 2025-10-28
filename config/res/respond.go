package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Respond envía una respuesta JSON uniforme
func Respond(c *gin.Context, statusCode int, status, message string, data interface{}, errors interface{}) {
	c.JSON(statusCode, gin.H{
		"status":  status, // "success" o "error"
		"code":    statusCode,
		"message": message,
		"data":    data,
		"errors":  errors,
	})
}

// Success envía respuesta exitosa
func Success(c *gin.Context, message string, data interface{}) {
	Respond(c, http.StatusOK, "success", message, data, nil)
}

// Created envía respuesta de creación
func Created(c *gin.Context, message string, data interface{}) {
	Respond(c, http.StatusCreated, "success", message, nil, nil)
}

// Error envía respuesta de error
func Error(c *gin.Context, statusCode int, message string, errors []string) {
	Respond(c, statusCode, "error", message, nil, errors)
}

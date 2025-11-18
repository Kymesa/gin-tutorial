package auth

import (
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"

	"gin-tutorial/internal/database"

	"gin-tutorial/config/jwt"
	"gin-tutorial/config/res"
)

func Test(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

func Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Email == "" || req.Password == "" {
		res.Error(c, http.StatusBadRequest, "Datos inválidos", nil)
		return
	}

	if len(req.Password) <= 2 {
		res.Error(c, http.StatusBadRequest, "Contraseña debe tener igual o mas de 3 caracteres", nil)
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := User{Email: req.Email, Password: string(hashed)}

	if err := database.DB.Create(&user).Error; err != nil {
		res.Error(c, http.StatusConflict, "Usuario ya existente", nil)
		return
	}
	res.Created(c, "Usuario registrado", nil)

}

func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	var user User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña incorrecta"})
		return
	}

	access, _ := jwt.GenerateJWT(user.ID)
	refresh, exp := jwt.RefreshJWT()

	rt := RefreshToken{UserID: user.ID, Token: refresh, ExpiresAt: exp}
	database.DB.Create(&rt)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Falta token"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.ValidateJWT(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Set("userId", claims.UserID)
		c.Next()
	}
}

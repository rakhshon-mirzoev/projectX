package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/internal/models"
	"main.go/internal/service"
)

func Login(c *gin.Context) {
	var input *models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка запроса"})
		return
	}
	id, err := service.CheckLoginPassword(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неправильно введен логин или пароль!"})
		return
	} else if id == 0 {
		c.JSON(http.StatusServiceUnavailable, "Просим пройти верификацию.")
		return
	}
	var token string
	if input.Type == "user" {
		token, err = GenerateToken(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка"})
			return
		}

	} else if input.Type == "university" {
		token, err = GenerateTokenForUnis(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка"})
			return
		}

	} else if input.Type == "school" {
		token, err = GenerateTokenForSchools(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка"})
			return
		}

	}
	c.JSON(http.StatusOK, gin.H{"message": "Доступ разрешен!", "token": token})

}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Нет доступа!")
			c.Abort()
			return
		}
		c.Next()
	}
}

func AdminLogin(c *gin.Context) {
	var login *models.LoginInput
	err := c.ShouldBindJSON(&login)
	if err != nil {
		return
	}
	adminId, err := service.CheckLoginAdmin(login)
	if err != nil {
		return
	}
	if adminId == 0 {
		c.JSON(http.StatusUnauthorized, "Нет доступа!")
		return
	}
	token, err := GenerateTokenForAdmins(adminId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Доступ разрешен.", "token": token})

}

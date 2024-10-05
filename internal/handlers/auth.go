package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"main.go/internal/models"
	"main.go/pkg/db"
)

func Login(c *gin.Context) {
	var input *models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка запроса"})
		return
	}
	id, err := checkLoginPassword(input)
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
		fmt.Println("user token,", token)
	} else if input.Type == "university" {
		token, err = GenerateTokenForUnis(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка"})
			return
		}
		fmt.Println("uni token", token)
	} else if input.Type == "school" {
		token, err = GenerateTokenForSchools(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка"})
			return
		}
		fmt.Println("school token", token)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Доступ разрешен!", "token": token})

}

func checkLoginPassword(input *models.LoginInput) (id int64, err error) {
	var existingUser models.User
	var existingUni models.University
	var existingSchool models.School
	if input.Type == "user" {
		err = db.GetDB().Raw("SELECT * FROM users WHERE login = ?", input.Login).Scan(&existingUser).Error
		if err != nil {
			return 0, err
		}
		fmt.Println("existing: ", existingUser)
		err = comparePasswordWithHash(input.Password, existingUser.Password)
		if err != nil {
			return 0, errors.New("пароль не совпадает с найденным")
		}
		return existingUser.Id, nil
	} else if input.Type == "university" {
		err = db.GetDB().Raw("SELECT * FROM university WHERE login = ?", input.Login).Scan(&existingUni).Error
		if err != nil {
			return 0, err
		}
		fmt.Println("existing: ", existingUni)
		err = comparePasswordWithHash(input.Password, existingUni.Password)
		if err != nil {
			return 0, errors.New("пароль не совпадает с найденным")
		}
		if !existingUni.Identification {
			return 0, nil
		}
		return existingUni.Id, nil
	} else if input.Type == "school" {
		err = db.GetDB().Raw("SELECT * FROM school WHERE login = ?", input.Login).Scan(&existingSchool).Error
		if err != nil {
			return 0, err
		}
		fmt.Println("existing: ", existingSchool)
		err = comparePasswordWithHash(input.Password, existingSchool.Password)
		if err != nil {
			return 0, errors.New("пароль не совпадает с найденным")
		}
		if !existingSchool.Identification {
			return 0, nil
		}
		return existingSchool.Id, nil
	} else {
		return 0, err
	}

}

func comparePasswordWithHash(input, hash string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(input))
	if err != nil {
		fmt.Println("input:", input, "hash:", hash)
		return err
	}
	fmt.Println("Password and Hash are identic.")
	return
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
		fmt.Println(err)
		return
	}
	adminId, err := checkLoginAdmin(login)
	if err != nil {
		fmt.Println(err)
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

func checkLoginAdmin(admin *models.LoginInput) (roleid int64, err error) {
	var existingAdmin models.User
	err = db.GetDB().Raw("SELECT * FROM users where login = ?", admin.Login).Scan(&existingAdmin).Error
	if err != nil {
		fmt.Println("Error in db")
		return 0, err
	}
	if existingAdmin.RoleId == 2 || existingAdmin.RoleId == 1 {
		return existingAdmin.Id, nil
	} else {
		return 0, nil
	}
}

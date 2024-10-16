package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"main.go/internal/models"
	"main.go/internal/service"
)

func CreateUser(c *gin.Context) {
	var user *models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Can't bind Your request")
		return
	}
	user.Password, err = hashString(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Can't bind Your request")
		return
	}
	err = service.CreateUser(user)
	if err != nil {
		if err.Error() == "user exists" {
			c.JSON(http.StatusBadRequest, "User with that login already exists")
			return
		} else if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, "Wrong param in Your request!")
			return
		} else {
			c.JSON(http.StatusInternalServerError, "Internal server error")
			return
		}
	}
	c.JSON(http.StatusOK, "Successfuly created!")
}

func DeleteUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}

	err = service.DeleteUserById(int64(userId))
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, "User with such {id} not found!")
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Successfuly deleted.": userId})
}

func GetUsers(c *gin.Context) {
	users, err := service.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}
	user, err := service.GetUserById(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	} else if user == nil {
		c.JSON(http.StatusBadRequest, "User with such {id} not found!")
		return
	}
	c.JSON(http.StatusOK, gin.H{"User": user})

}

func hashString(input string) (hashedPassword string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

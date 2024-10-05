package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/internal/models"
	"main.go/internal/service"
)

func CreateSocialScien(c *gin.Context) {
	var socialscien *models.SocialScien
	if err := c.ShouldBindJSON(&socialscien); err != nil {
		c.JSON(http.StatusBadRequest, "Can't bind your request")
		return
	}
	err := service.CreateSocialScien(socialscien)
	if err != nil {
		if err.Error() == "socialscien exists" {
			c.JSON(http.StatusBadRequest, "Socialscien with that name already exists")
			return
		} else if err.Error() == "record not found" {
			c.JSON(http.StatusOK, gin.H{"Created an socialscien": socialscien})
		} else {
			c.JSON(http.StatusInternalServerError, "Internal server error")
			return
		}
	}

}

func GetSocialSciences(c *gin.Context) {
	socialsciences, err := service.GetSocialSciences()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Socialsciences:": socialsciences})
}

func GetSocialScienById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}
	socialscien, err := service.GetSocialScienById(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	} else if socialscien == nil {
		c.JSON(http.StatusBadRequest, "Socialscien with such {id} not found!")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Socialscien": socialscien})
}

func DeleteSocialScienById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}
	err = service.DeleteSocialScienById(int64(id))
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, "Socialscien with such {id} not found!")
		return

	} else if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Successfuly deleted.": id})
}

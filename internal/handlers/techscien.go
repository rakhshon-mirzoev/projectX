package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/internal/models"
	"main.go/internal/service"
)

func CreateTechScien(c *gin.Context) {
	var techscien *models.TechScien
	if err := c.ShouldBindJSON(&techscien); err != nil {
		c.JSON(http.StatusBadRequest, "Can't bind your request")
		return
	}
	err := service.CreateTechScien(techscien)
	if err != nil {
		if err.Error() == "techscien exists" {
			c.JSON(http.StatusBadRequest, "Techscien with that name already exists")
			return
		} else if err.Error() == "record not found" {
			c.JSON(http.StatusOK, gin.H{"Created an techscien": techscien})
		} else {
			c.JSON(http.StatusInternalServerError, "Internal server error")
			return
		}
	}
}

func GetTechSciences(c *gin.Context) {
	techsciences, err := service.GetTechSciences()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Techsciences:": techsciences})
}

func GetTechScienById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}
	techscien, err := service.GetTechScienById(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	} else if techscien == nil {
		c.JSON(http.StatusBadRequest, "Techscien with such {id} not found!")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Techscien": techscien})
}

func DeleteTechScienById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}
	err = service.DeleteTechScienById(int64(id))
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, "Techscien with such {id} not found!")
		return

	} else if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Successfuly deleted.": id})
}

package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/internal/models"
	"main.go/internal/service"
)

func CreateHumanScien(c *gin.Context) {
	var humanscien *models.HumanScien
	if err := c.ShouldBindJSON(&humanscien); err != nil {
		c.JSON(http.StatusBadRequest, "Can't bind your request")
		return
	}
	err := service.CreateHumanScien(humanscien)
	if err != nil {
		if err.Error() == "humanscien exists" {
			c.JSON(http.StatusBadRequest, "Humanscien with that name already exists")
			return
		} else if err.Error() == "record not found" {
			c.JSON(http.StatusOK, gin.H{"Created an humanscien": humanscien})
		} else {
			c.JSON(http.StatusInternalServerError, "Internal server error")
			return
		}
	}

}

func GetHumanSciences(c *gin.Context) {
	humansciences, err := service.GetHumanSciences()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Humansciences:": humansciences})
}

func GetHumanScienById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}
	humanscien, err := service.GetHumanScienById(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	} else if humanscien == nil {
		c.JSON(http.StatusBadRequest, "Humanscien with such {id} not found!")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Humanscien": humanscien})
}

func DeleteHumanScienById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}
	err = service.DeleteHumanScienById(int64(id))
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, "Humanscien with such {id} not found!")
		return

	} else if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Successfuly deleted.": id})
}

package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/internal/models"
	"main.go/internal/service"
)

func CreateExactScien(c *gin.Context) {
	var exactscien *models.ExactScien
	if err := c.ShouldBindJSON(&exactscien); err != nil {
		c.JSON(http.StatusBadRequest, "Can't bind your request")
		return
	}
	err := service.CreateExactScien(exactscien)
	if err != nil {
		if err.Error() == "exactscien exists" {
			c.JSON(http.StatusBadRequest, "Exactscien with that name already exists")
			return
		} else if err.Error() == "record not found" {
			c.JSON(http.StatusOK, gin.H{"Created an exactscien": exactscien})
		} else {
			c.JSON(http.StatusInternalServerError, "Internal server error")
			return
		}
	}

}

func GetExactSciences(c *gin.Context) {
	exactsciences, err := service.GetExactSciences()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Exactsciences:": exactsciences})
}

func GetExactScienById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}
	exactscien, err := service.GetExactScienById(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	} else if exactscien == nil {
		c.JSON(http.StatusBadRequest, "Exactscien with such {id} not found!")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Exactscien": exactscien})
}

func DeleteExactScienById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}
	err = service.DeleteExactScienById(int64(id))
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, "Exactscien with such {id} not found!")
		return

	} else if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Successfuly deleted.": id})
}

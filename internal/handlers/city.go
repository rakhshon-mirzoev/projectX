package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/internal/models"
	"main.go/internal/service"
)

func CreateCity(c *gin.Context) {
	var city *models.City
	if err := c.ShouldBindJSON(&city); err != nil {
		c.JSON(http.StatusBadRequest, "Can't bind your request")
		return
	}
	err := service.CreateCity(city)
	if err != nil {
		if err.Error() == "city exists" {
			c.JSON(http.StatusBadRequest, "City with that name already exists")
			return
		} else if err.Error() == "record not found" {
			c.JSON(http.StatusOK, gin.H{"Created a city": city})
		} else {
			c.JSON(http.StatusInternalServerError, "Internal server error")
			return
		}
	}

}

func GetCities(c *gin.Context) {
	cities, err := service.GetCities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Cities:": cities})
}

func GetCityById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}
	city, err := service.GetCityById(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	} else if city == nil {
		c.JSON(http.StatusBadRequest, "City with such {id} not found!")
		return
	}
	c.JSON(http.StatusOK, gin.H{"City": city})
}

func DeleteCityById(c *gin.Context) {
	cityId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}
	err = service.DeleteCityById(int64(cityId))
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, "City with such {id} not found!")
		return

	} else if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Successfuly deleted.": cityId})
}

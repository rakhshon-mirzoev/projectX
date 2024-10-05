package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/internal/models"
	"main.go/internal/service"
)

func CreateSchool(c *gin.Context) {
	var school *models.School
	err := c.BindJSON(&school)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Can't bind your request")
		return
	}
	school.Password, err = hashString(school.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Can't bind Your request")
		return
	}
	err = service.CreateSchool(school)
	if err != nil {
		if err.Error() == "school exists" {
			c.JSON(http.StatusBadRequest, "School with that name already exists")
			return
		} else {
			c.JSON(http.StatusInternalServerError, "Internal server error")
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"Created a school": school})
}

func DeleteSchoolById(c *gin.Context) {
	schId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}

	err = service.DeleteSchoolById(int64(schId))
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, "School with such {id} not found!")
		return

	} else if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Successfuly deleted.": schId})
}

func GetSchools(c *gin.Context) {
	var f models.SchoolFilter
	err := c.Bind(&f)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Can't bind your filter!")
		return
	}

	schools, err := service.GetSchools(&f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Schools:": schools})
}

func GetSchoolById(c *gin.Context) {
	schId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}
	school, err := service.GetSchoolById(int64(schId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	} else if school == nil {
		c.JSON(http.StatusBadRequest, "School with such {id} not found!")
		return
	}
	c.JSON(http.StatusOK, gin.H{"School:": school})
}

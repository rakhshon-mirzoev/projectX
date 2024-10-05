package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/internal/models"
	"main.go/internal/service"
)

func CreateUni(c *gin.Context) {
	var uni *models.University
	if err := c.ShouldBindJSON(&uni); err != nil {
		c.JSON(http.StatusBadRequest, "Can't bind your request")
		return
	}
	var err error
	uni.Password, err = hashString(uni.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Can't bind Your request")
		return
	}
	err = service.CreateUni(uni)
	if err != nil {
		if err.Error() == "university exists" {
			c.JSON(http.StatusBadRequest, "University with that name already exists")
			return
		} else if err.Error() == "record not found" {
			c.JSON(http.StatusOK, gin.H{"Created University with data:": uni})
		} else {
			c.JSON(http.StatusInternalServerError, "Internal server error")
			return
		}
	}

}

func DeleteUniById(c *gin.Context) {
	uniId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}

	err = service.DeleteUniById(int64(uniId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, "Uni with such {id} not found!")
			return
		} else {
			c.JSON(http.StatusInternalServerError, "Internal server error(")
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"Successfuly deleted.": uniId})
}

func GetUnis(c *gin.Context) {
	var f models.UniversityFilter
	err := c.Bind(&f)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Can't bind your filter!")
		return
	}
	unis, err := service.GetUnis(&f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Universities": unis})
}

func GetUniById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}
	uni, err := service.GetUniById(int64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, "University with such {id} not found")
			return
		} else {
			c.JSON(http.StatusInternalServerError, "Internal server error(")
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"University": uni})
}

func GetUniExactSubs(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}
	exsubs, err := service.GetUniExactSubs(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	} else if exsubs == nil {
		c.JSON(http.StatusNotFound, "Exact sciences not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{"University exact sciences": exsubs})
}

func GetUniHumanSubs(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}
	husubs, err := service.GetUniHumanSubs(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"University human sciences": husubs})
}

func GetUniTechSubs(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}
	tesubs, err := service.GetUniTechSubs(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"University tech sciences": tesubs})
}

func GetUniSocialSubs(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}
	sosubs, err := service.GetUniSocialSubs(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, gin.H{"University social sciences": sosubs})
}
func GetUniWithCityName(c *gin.Context) {
	universitiesWithCityNames, err := service.GetUniWithCityName()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	}
	c.JSON(http.StatusOK, universitiesWithCityNames)
}

func ChangeUniversity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot find param {id}!")
		return
	}
	var f models.UniversityUpdateFilter
	err = c.Bind(&f)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Can't bind your filter!")
		return
	}
	uni, err := service.ChangeUniversity(int64(id), f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error(")
		return
	} else if uni == nil {
		c.JSON(http.StatusBadRequest, "University with such {id} not found!")
		return
	}
	c.JSON(http.StatusOK, uni)
}

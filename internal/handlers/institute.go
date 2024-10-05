package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/internal/models"
	"main.go/internal/service"
)

func GetInstitutionList(c *gin.Context) {
	var (
		f  models.Institute
		fU models.UniversityFilter
		fS models.SchoolFilter
	)
	err := c.Bind(&f)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Can't bind type")
		return
	}
	err = c.Bind(&fU)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Can't bind type")
		return
	}
	err = c.Bind(&fS)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Can't bind type")
		return
	}
	sc, u, err := service.GetInstitutionList(f, &fU, &fS)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Schools:": sc, "Unis:": u})
}

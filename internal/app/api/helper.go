package api

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"main.go/internal/handlers"
	"main.go/pkg/constants"
)

// Пытаемся отконфигурировать наш API_INSTANCE (а конкретнее поле logger)
func (a *API) configureLoggerField() error {
	info_level, err := logrus.ParseLevel(a.config.InfoLevel)
	if err != nil {
		return err
	}
	// Подтверждение того, что логгер сконфигурировался
	a.logger.SetLevel(info_level)
	return nil
}

// Пытаемся сконфигурировать маршрутизатор (а конкретнее поле router API)
func (a *API) configureRouterField() {
	//Cors
	corsMiddleWare := CorsMiddleWare()

	//Open points
	a.router.POST("/Login", handlers.Login)
	a.router.POST("/UserRegistration", handlers.CreateUser)
	a.router.POST("/UniversityRegistration", handlers.CreateUni)
	a.router.POST("/SchoolRegistration", handlers.CreateSchool)
	a.router.GET("/GetInstitutionList", handlers.GetInstitutionList)
	a.router.GET("/GetUniversitiesWithCityName", handlers.GetUniWithCityName)
	a.router.POST("/Admin", handlers.AdminLogin)
	a.router.GET("/ping", PingPong)
	//Closed endpoints
	cl := a.router.Group(constants.ApiPrefix)
	cl.Use(handlers.JwtAuthMiddleware(), corsMiddleWare)
	{
		// Users
		{
			cl.POST("/CreateUser", handlers.CreateUser)
			cl.DELETE("/DeleteUser/:id", handlers.DeleteUserById)
			cl.GET("/GetUsers", handlers.GetUsers)
			cl.GET("/GetUserById/:id", handlers.GetUserById)
		}

		// School
		{
			cl.POST("/CreateSchool", handlers.CreateSchool)
			cl.GET("/GetSchools", handlers.GetSchools)
			cl.GET("/GetSchoolById/:id", handlers.GetSchoolById)
			cl.DELETE("/DeleteSchool/:id", handlers.DeleteSchoolById)
		}

		// University
		{
			cl.POST("/CreateUniversity", handlers.CreateUni)
			cl.DELETE("/DeleteUniversity/:id", handlers.DeleteUniById)
			cl.GET("/GetUniversities", handlers.GetUnis)
			cl.GET("/GetUniversityById/:id", handlers.GetUniById)
			cl.GET("/GetUniExactSubs/:id", handlers.GetUniExactSubs)
			cl.GET("/GetUniHumanSubs/:id", handlers.GetUniHumanSubs)
			cl.GET("/GetUniTechSubs/:id", handlers.GetUniTechSubs)
			cl.GET("/GetUniSocialSubs/:id", handlers.GetUniSocialSubs)
			cl.PUT("/ChangeUniversity/:id", handlers.ChangeUniversity)
		}

		// City
		{
			cl.POST("/CreateCity", handlers.CreateCity)
			cl.DELETE("/DeleteCity/:id", handlers.DeleteCityById)
			cl.GET("/GetCities", handlers.GetCities)
			cl.GET("/GetCityById/:id", handlers.GetCityById)
		}

		// ExactScience
		{
			cl.POST("/CreateExactScience", handlers.CreateExactScien)
			cl.DELETE("/DeleteExactScience/:id", handlers.DeleteExactScienById)
			cl.GET("/GetExactSciences", handlers.GetExactSciences)
			cl.GET("/GetExactScienceById/:id", handlers.GetExactScienById)
		}

		// HumanScience
		{
			cl.POST("/CreateHumanScience", handlers.CreateHumanScien)
			cl.DELETE("/DeleteHumanScience/:id", handlers.DeleteHumanScienById)
			cl.GET("/GetHumanSciences", handlers.GetHumanSciences)
			cl.GET("/GetHumanScienceById/:id", handlers.GetHumanScienById)
		}

		// TechScience
		{
			cl.POST("/CreateTechScience", handlers.CreateTechScien)
			cl.DELETE("/DeleteTechScience/:id", handlers.DeleteTechScienById)
			cl.GET("/GetTechSciences", handlers.GetTechSciences)
			cl.GET("/GetTechScienceById/:id", handlers.GetTechScienById)
		}

		// SocialScien
		{
			cl.POST("/CreateSocialScience", handlers.CreateSocialScien)
			cl.DELETE("/DeleteSocialScience/:id", handlers.DeleteSocialScienById)
			cl.GET("/GetSocialSciences", handlers.GetSocialSciences)
			cl.GET("/GetSocialScienceById/:id", handlers.GetSocialScienById)
		}
	}
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})

}

func CorsMiddleWare() gin.HandlerFunc {
	config := cors.DefaultConfig()

	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	a := cors.New(config)
	return a
}

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"main.go/pkg/db"
)

type API struct {
	config   *Config
	logger   *logrus.Logger
	router   *gin.Engine
	database *gorm.DB
}

func New(config *Config) *API {
	api := &API{
		config: config,
		logger: logrus.New(),
		router: gin.New(),
	}
	api.openDB()
	return api
}

func (api *API) openDB() {
	db.StartDatabase()
	api.database = db.GetDB()
}

func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	api.logger.Info("Starting api server at port", api.config.Port)

	api.configureRouterField()
	// На этапе валидного завершения запускаем сервер
	return http.ListenAndServe(api.config.Port, api.router)
}

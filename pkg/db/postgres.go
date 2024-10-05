package db

import (
	log "main.go/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"main.go/internal/models"
	"main.go/pkg/constants"
)

var dbInstance *gorm.DB

func StartDatabase() {
	var err error
	dbInstance, err = gorm.Open(postgres.Open(constants.DbCredentials), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Error.Fatalf("db.setupPostgresProduction err: %v", err)
	}

	sqlDB, err := dbInstance.DB()
	if err != nil {
		log.Error.Fatalf("Error on getting the DB instance: %v", err)
	}
	sqlDB.SetMaxOpenConns(100)

	if err := autoMigrate(); err != nil {
		log.Error.Fatalf("Error during migration: %v", err)
	}

	log.Info.Println("DB successfully connected!")
}

func autoMigrate() error {
	for _, model := range []interface{}{
		&models.User{},
		&models.Role{},
		&models.School{},
		&models.University{},
		&models.City{},
		&models.HumanScien{},
		&models.ExactScien{},
		&models.SocialScien{},
		&models.TechScien{},
		&models.Faculty{},
		&models.UniversityExacts{},
		&models.UniversitySocials{},
		&models.UniversityTechs{},
		&models.UniversityHumans{},
		&models.UniversityFaculty{},
		&models.FavouriteUnis{},
	} {
		if err := dbInstance.AutoMigrate(model); err != nil {
			log.Error.Fatalf("auto migrate %T: %T", model, err)
			return err
		}
	}
	return nil
}

func GetDB() *gorm.DB {
	return dbInstance
}

func CloseDB() {
	sqlDB, err := dbInstance.DB()
	if err != nil {
		log.Error.Fatalf("Error on getting the DB instance: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Error.Fatalf("Error on closing the DB: %v", err)
	}
	log.Info.Println("Db closed successfully")
}

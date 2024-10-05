package service

import (
	"main.go/internal/models"
	"main.go/internal/repository"
	log "main.go/logger"
)

func GetCities() (cities []*models.City, err error) {
	cities, err = repository.GetCities()
	if err != nil {
		log.Error.Println("Error in repository (GetCities): ", err)
		return cities, err
	}
	return cities, err
}

func DeleteCityById(cityId int64) error {
	err := repository.DeleteCityById(cityId)
	if err != nil {
		log.Error.Println("Error in repository (DeleteCityById): ", err)
		return err
	}
	return nil
}

func CreateCity(city *models.City) error {
	err := repository.CreateCity(city)
	if err != nil {
		if err.Error() == "city exists" {
			log.Warn.Println("City with that name already exists!")
			return err
		} else if err.Error() == "record not found" {
			return nil
		} else {
			log.Error.Println("Error in repository (CreateCity):", err)
			return err
		}
	} else {
		log.Info.Println("Successfuly created a new city.")
	}
	return nil
}

func GetCityById(cityId int64) (city *models.City, err error) {
	city, err = repository.GetCityById(cityId)
	if err != nil {
		log.Error.Println("Error in repository (GetCityById):", err)
		return nil, err
	}
	return city, nil
}

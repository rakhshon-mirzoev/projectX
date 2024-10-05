package service

import (
	log "main.go/logger"

	"main.go/internal/models"
	"main.go/internal/repository"
)

func CreateExactScien(exactscien *models.ExactScien) error {
	err := repository.CreateExactScien(exactscien)
	if err != nil {
		if err.Error() == "exactscien exists" {
			log.Warn.Println("Exactscien with that name already exists!")
			return err
		} else if err.Error() == "record not found" {
			return nil
		} else {
			log.Error.Println("Error in repository (CreateExactScien):", err)
			return err
		}
	} else {
		log.Info.Println("Successfuly created a new exactscien.")
	}
	return nil
}

func DeleteExactScienById(exactscienId int64) error {
	err := repository.DeleteExactScienById(exactscienId)
	if err != nil {
		log.Error.Println("Error in repository (DeleteExactScienById): ", err)
		return err
	}
	return nil
}

func GetExactSciences() (exactsciences []*models.ExactScien, err error) {
	exactsciences, err = repository.GetExactSciences()
	if err != nil {
		log.Error.Println("Error in repository (GetExactSciences): ", err)
		return nil, err
	}
	return exactsciences, err
}

func GetExactScienById(exactscienId int64) (exactscien *models.ExactScien, err error) {
	exactscien, err = repository.GetExactScienById(exactscienId)
	if err != nil {
		log.Error.Println("Error in repository (GetExactScienById):", err)
		return nil, err
	}
	return exactscien, nil
}

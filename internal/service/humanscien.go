package service

import (
	log "main.go/logger"

	"main.go/internal/models"
	"main.go/internal/repository"
)

func CreateHumanScien(humanscien *models.HumanScien) error {
	err := repository.CreateHumanScien(humanscien)
	if err != nil {
		if err.Error() == "humanscien exists" {
			log.Warn.Println("Humanscien with that name already exists!")
			return err
		} else if err.Error() == "record not found" {
			return nil
		} else {
			log.Error.Println("Error in repository (CreateHumanScienScien):", err)
			return err
		}
	} else {
		log.Info.Println("Successfuly created a new humanscien.")
	}
	return nil
}

func DeleteHumanScienById(humanscienId int64) error {
	err := repository.DeleteHumanScienById(humanscienId)
	if err != nil {
		log.Error.Println("Error in repository (DeleteHumanScienById): ", err)
		return err
	}
	return nil
}

func GetHumanSciences() (humansciences []*models.HumanScien, err error) {
	humansciences, err = repository.GetHumanSciences()
	if err != nil {
		log.Error.Println("Error in repository (GetHumanSciences): ", err)
		return nil, err
	}
	return humansciences, err
}

func GetHumanScienById(humanscienId int64) (humanscien *models.HumanScien, err error) {
	humanscien, err = repository.GetHumanScienById(humanscienId)
	if err != nil {
		log.Error.Println("Error in repository (GetHumanScienById):", err)
		return nil, err
	}
	return humanscien, nil
}

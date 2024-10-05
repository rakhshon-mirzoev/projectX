package service

import (
	log "main.go/logger"

	"main.go/internal/models"
	"main.go/internal/repository"
)

func CreateTechScien(techscien *models.TechScien) error {
	err := repository.CreateTechScien(techscien)
	if err != nil {
		if err.Error() == "techscien exists" {
			log.Warn.Println("Techscien with that name already exists!")
			return err
		} else if err.Error() == "record not found" {
			return nil
		} else {
			log.Error.Println("Error in repository (CreateTechScienScien):", err)
			return err
		}
	} else {
		log.Info.Println("Successfuly created a new techscien.")
	}
	return nil
}

func DeleteTechScienById(techscienId int64) error {
	err := repository.DeleteTechScienById(techscienId)
	if err != nil {
		log.Error.Println("Error in repository (DeleteTechScienById): ", err)
		return err
	}
	return nil
}

func GetTechSciences() (techsciences []*models.TechScien, err error) {
	techsciences, err = repository.GetTechSciences()
	if err != nil {
		log.Error.Println("Error in repository (GetTechSciences): ", err)
		return nil, err
	}
	return techsciences, err
}

func GetTechScienById(humanscienId int64) (techscien *models.TechScien, err error) {
	techscien, err = repository.GetTechScienById(humanscienId)
	if err != nil {
		log.Error.Println("Error in repository (GetTechScienById):", err)
		return nil, err
	}
	return techscien, nil
}

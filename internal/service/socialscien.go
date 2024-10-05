package service

import (
	log "main.go/logger"

	"main.go/internal/models"
	"main.go/internal/repository"
)

func CreateSocialScien(socialscien *models.SocialScien) error {
	err := repository.CreateSocialScien(socialscien)
	if err != nil {
		if err.Error() == "socialscien exists" {
			log.Warn.Println("Socialscien with that name already exists!")
			return err
		} else if err.Error() == "record not found" {
			return nil
		} else {
			log.Error.Println("Error in repository (CreateSocialScienScien):", err)
			return err
		}
	} else {
		log.Info.Println("Successfuly created a new socialscien.")
	}
	return nil
}

func DeleteSocialScienById(socialscienId int64) error {
	err := repository.DeleteSocialScienById(socialscienId)
	if err != nil {
		log.Error.Println("Error in repository (DeleteSocialScienById): ", err)
		return err
	}
	return nil
}

func GetSocialSciences() (socialsciences []*models.SocialScien, err error) {
	socialsciences, err = repository.GetSocialSciences()
	if err != nil {
		log.Error.Println("Error in repository (GetSocialSciences): ", err)
		return nil, err
	}
	return socialsciences, err
}

func GetSocialScienById(socialscienId int64) (socialscien *models.SocialScien, err error) {
	socialscien, err = repository.GetSocialScienById(socialscienId)
	if err != nil {
		log.Error.Println("Error in repository (GetSocialScienById):", err)
		return nil, err
	}
	return socialscien, nil
}

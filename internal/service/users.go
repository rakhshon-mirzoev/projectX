package service

import (
	"main.go/internal/models"
	"main.go/internal/repository"
	log "main.go/logger"
)

func CreateUser(user *models.User) error {
	err := repository.CreateUser(user)
	if err != nil {
		if err.Error() == "user exists" {
			log.Warn.Println("User with that login already exists!")
			return err
		} else if err.Error() == "record not found" {
			return nil
		} else {
			log.Error.Fatal("Error in repository(CreateUser):", err)
			return err
		}
	} else {
		log.Info.Println("Successfully created a new user.")
	}
	return nil
}

func DeleteUserById(userId int64) error {
	err := repository.DeleteUserById(userId)
	if err != nil {
		log.Error.Fatal("Error in repository (DeleteUserById):", err)
		return err
	}
	return nil
}

func GetUsers() (u *[]models.User, err error) {
	u, err = repository.GetUsers()
	if err != nil {
		log.Error.Fatal("Error in repo GetUsers", err)
		return nil, err
	}
	return u, nil
}

func GetUserById(userId int64) (u *models.User, err error) {
	u, err = repository.GetUserById(userId)
	if err != nil {
		log.Error.Fatal("Error in GetUserById", err)
		return nil, err
	}
	return u, nil
}

package service

import (
	"main.go/internal/models"
	"main.go/internal/repository"
	"main.go/logger"
)

func CheckLoginAdmin(admin *models.LoginInput) (roleid int64, err error) {
	roleid, err = repository.CheckLoginAdmin(admin)
	if err != nil {
		logger.Error.Println("Error in repository (CheckLoginAdmin)", err)
		return 0, err
	}
	return roleid, err
}

func CheckLoginPassword(input *models.LoginInput) (id int64, err error) {
	id, err = repository.CheckLoginPassword(input)
	if err != nil {
		return 0, err
	}
	return id, err
}

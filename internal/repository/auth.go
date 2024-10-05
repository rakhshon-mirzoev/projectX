package repository

import (
	"main.go/internal/models"
	"main.go/internal/token"
	"main.go/pkg/db"
)

var (
	existingUser   models.User
	existingUni    models.University
	existingSchool models.School
)

func CheckLoginAdmin(admin *models.LoginInput) (roleid int64, err error) {
	var existingAdmin models.User
	err = db.GetDB().Raw("SELECT * FROM users where login = ?", admin.Login).Scan(&existingAdmin).Error
	if err != nil {
		return 0, err
	}
	if existingAdmin.RoleId == 2 || existingAdmin.RoleId == 1 {
		return existingAdmin.Id, nil
	} else {
		return 0, nil
	}
}

func CheckLoginPassword(input *models.LoginInput) (id int64, err error) {
	if input.Type == "user" {
		err = db.GetDB().Raw("SELECT * FROM users WHERE login = ?", input.Login).Scan(&existingUser).Error

		if err != nil {
			return 0, err
		}

		err = token.ComparePasswordWithHash(input.Password, existingUser.Password)

		if err != nil {
			return 0, err
		}
		return existingUser.Id, nil

	} else if input.Type == "university" {
		err = db.GetDB().Raw("SELECT * FROM university WHERE login = ?", input.Login).Scan(&existingUni).Error
		if err != nil {
			return 0, err
		}

		err = token.ComparePasswordWithHash(input.Password, existingUni.Password)

		if err != nil {
			return 0, err
		}
		if !existingUni.Identification {
			return 0, nil
		}
		return existingUni.Id, nil

	} else if input.Type == "school" {
		err = db.GetDB().Raw("SELECT * FROM school WHERE login = ?", input.Login).Scan(&existingSchool).Error

		if err != nil {
			return 0, err
		}

		err = token.ComparePasswordWithHash(input.Password, existingSchool.Password)

		if err != nil {
			return 0, err
		}
		if !existingSchool.Identification {
			return 0, nil
		}
		return existingSchool.Id, nil

	} else {
		return 0, err
	}

}

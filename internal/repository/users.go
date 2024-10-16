package repository

import (
	"errors"

	"gorm.io/gorm"
	"main.go/internal/models"
	"main.go/pkg/db"
)

func CreateUser(user *models.User) error {
	var cityId int64
	cityId, err := FindCityId(user.CityName)
	if err != nil {
		return translateError(err)
	} else if cityId == 0 {
		return gorm.ErrRecordNotFound
	}
	existing, err := userExists(user.Login)
	if err != nil {
		return err
	}
	if existing {
		return errors.New("user exists")
	}
	if user.UserType == "Студент/Студентка" {

		sqlReq := `INSERT INTO users(name, phone, age, city_id, login, password, role_id) VALUES(?,?,?,?,?,?,4)`
		err = db.GetDB().Exec(sqlReq, user.Name, user.Phone, user.Age, cityId, user.Login, user.Password).Error
		if err != nil {
			return err
		}
	} else if user.UserType == "Школьник/Школьница" {
		sqlReq := `INSERT INTO users(name, phone, age, city_id, login, password, role_id) VALUES(?,?,?,?,?,?,5)`
		err = db.GetDB().Exec(sqlReq, user.Name, user.Phone, user.Age, cityId, user.Login, user.Password).Error
		if err != nil {
			return err
		}
	} else if user.UserType == "Родитель" {
		sqlReq := `INSERT INTO users(name, phone, age, city_id, login, password, role_id) VALUES(?,?,?,?,?,?,8)`
		err = db.GetDB().Exec(sqlReq, user.Name, user.Phone, user.Age, cityId, user.Login, user.Password).Error
		if err != nil {
			return err
		}
	} else {
		sqlReq := `INSERT INTO users(name, phone, age, city_id, login, password, role_id) VALUES(?,?,?,?,?,?,?)`
		err = db.GetDB().Exec(sqlReq, user.Name, user.Phone, user.Age, cityId, user.Login, user.Password, user.RoleId).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteUserById(userId int64) error {
	var u *models.User
	if err := db.GetDB().First(&u, userId).Error; err != nil {
		return translateError(err)
	}
	err := db.GetDB().Delete(&u, userId).Error
	if err != nil {
		return translateError(err)
	}
	return nil
}

func GetUsers() (u *[]models.User, err error) {
	sqlReq := `SELECT * FROM users ORDER BY name`
	err = db.GetDB().Raw(sqlReq).Scan(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserById(userId int64) (u *models.User, err error) {
	sqlReq := `select * from users where id = ?`
	err = db.GetDB().Raw(sqlReq, userId).Scan(&u).Error
	if err != nil {
		return u, err
	}
	return u, nil
}

func userExists(login string) (bool, error) {
	var user models.User
	sqlReq := `SELECT * FROM users WHERE login = ?`
	err := db.GetDB().Raw(sqlReq, login).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

package repository

import (
	"errors"

	"gorm.io/gorm"
	"main.go/internal/models"
	"main.go/pkg/db"
)

func CreateCity(city *models.City) error {
	existing, err := cityExists(city.Name)
	if err != nil {
		return err
	}
	if existing {
		return errors.New("city exists")
	}
	sqlReq := `INSERT INTO city(name) VALUES(?)`
	err = db.GetDB().Exec(sqlReq, city.Name).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteCityById(cityId int64) error {
	var city *models.City
	if err := db.GetDB().First(&city, cityId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	err := db.GetDB().Delete(&city, cityId).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCities() (cities []*models.City, err error) {
	sqlReq := `SELECT * FROM city ORDER BY name`
	err = db.GetDB().Raw(sqlReq).Scan(&cities).Error
	if err != nil {
		return cities, err
	}
	return cities, err
}

func GetCityById(cityId int64) (city *models.City, err error) {
	sqlReq := `SELECT * FROM city WHERE id = ?`
	err = db.GetDB().Raw(sqlReq, cityId).Scan(&city).Error
	if err != nil {
		return city, err
	}
	return city, nil
}

func cityExists(name string) (bool, error) {
	var city models.City
	sqlReq := `SELECT * FROM city WHERE name = ?`
	err := db.GetDB().Raw(sqlReq, name).First(&city).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func FindCityId(name string) (id int64, err error) {
	var city models.City
	err = db.GetDB().Table("city").Where("name = ?", name).Find(&city).Error
	if err != nil {
		return 0, err
	}
	return city.Id, nil
}

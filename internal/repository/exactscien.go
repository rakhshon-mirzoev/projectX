package repository

import (
	"errors"

	"gorm.io/gorm"
	"main.go/internal/models"
	"main.go/pkg/db"
)

func CreateExactScien(sub *models.ExactScien) error {
	existing, err := exactScienExists(sub.Name)
	if err != nil {
		return err
	}
	if existing {
		return errors.New("exactscien exists")
	}
	sqlReq := `INSERT INTO exactscien(name) VALUES(?)`
	err = db.GetDB().Exec(sqlReq, sub.Name).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteExactScienById(exactscienId int64) error {
	var exactscien *models.ExactScien
	if err := db.GetDB().First(&exactscien, exactscienId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	err := db.GetDB().Delete(&exactscien, exactscienId).Error
	if err != nil {
		return err
	}
	return nil
}

func GetExactSciences() (exactsciences []*models.ExactScien, err error) {
	sqlReq := `SELECT * FROM exactscien ORDER BY name`
	err = db.GetDB().Raw(sqlReq).Scan(&exactsciences).Error
	if err != nil {
		return nil, err
	}
	return exactsciences, err
}

func GetExactScienById(exactscienId int64) (exactscien *models.ExactScien, err error) {
	sqlReq := `SELECT * FROM exactscien WHERE id = ?`
	err = db.GetDB().Raw(sqlReq, exactscienId).Scan(&exactscien).Error
	if err != nil {
		return nil, err
	}
	return exactscien, nil
}

func exactScienExists(name string) (bool, error) {
	var exactscien *models.ExactScien
	sqlReq := `SELECT * FROM exactscien WHERE name = ?`
	err := db.GetDB().Raw(sqlReq, name).First(&exactscien).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

package repository

import (
	"errors"

	"gorm.io/gorm"
	"main.go/internal/models"
	"main.go/pkg/db"
)

func CreateHumanScien(humanscien *models.HumanScien) error {
	existing, err := humanScienExists(humanscien.Name)
	if err != nil {
		return err
	}
	if existing {
		return errors.New("humanscien exists")
	}
	sqlReq := `INSERT INTO humanscien(name) VALUES(?)`
	err = db.GetDB().Exec(sqlReq, humanscien.Name).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteHumanScienById(humanscienId int64) error {
	var humanscien *models.HumanScien
	if err := db.GetDB().First(&humanscien, humanscienId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	err := db.GetDB().Delete(&humanscien, humanscienId).Error
	if err != nil {
		return err
	}
	return nil
}

func GetHumanSciences() (humansciences []*models.HumanScien, err error) {
	sqlReq := `SELECT * FROM humanscien ORDER BY name`
	err = db.GetDB().Raw(sqlReq).Scan(&humansciences).Error
	if err != nil {
		return nil, err
	}
	return humansciences, err
}

func GetHumanScienById(humanscienId int64) (humanscien *models.HumanScien, err error) {
	sqlReq := `SELECT * FROM humanscien WHERE id = ?`
	err = db.GetDB().Raw(sqlReq, humanscienId).Scan(&humanscien).Error
	if err != nil {
		return nil, err
	}
	return humanscien, nil
}

func humanScienExists(name string) (bool, error) {
	var humanscien *models.HumanScien
	sqlReq := `SELECT * FROM humanscien WHERE name = ?`
	err := db.GetDB().Raw(sqlReq, name).First(&humanscien).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

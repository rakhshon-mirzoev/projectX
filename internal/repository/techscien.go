package repository

import (
	"errors"

	"gorm.io/gorm"
	"main.go/internal/models"
	"main.go/pkg/db"
)

func CreateTechScien(techscien *models.TechScien) error {
	existing, err := techScienExists(techscien.Name)
	if err != nil {
		return err
	}
	if existing {
		return errors.New("techscien exists")
	}
	sqlReq := `INSERT INTO techscien(name) VALUES(?)`
	err = db.GetDB().Exec(sqlReq, techscien.Name).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteTechScienById(techscienId int64) error {
	var techscien *models.TechScien
	if err := db.GetDB().First(&techscien, techscienId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	err := db.GetDB().Delete(&techscien, techscienId).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTechSciences() (techsciences []*models.TechScien, err error) {
	sqlReq := `SELECT * FROM techscien ORDER BY name`
	err = db.GetDB().Raw(sqlReq).Scan(&techsciences).Error
	if err != nil {
		return nil, err
	}
	return techsciences, err
}

func GetTechScienById(techscienId int64) (techscien *models.TechScien, err error) {
	sqlReq := `SELECT * FROM techscien WHERE id = ?`
	err = db.GetDB().Raw(sqlReq, techscienId).Scan(&techscien).Error
	if err != nil {
		return nil, err
	}
	return techscien, nil
}

func techScienExists(name string) (bool, error) {
	var techscien *models.TechScien
	sqlReq := `SELECT * FROM techscien WHERE name = ?`
	err := db.GetDB().Raw(sqlReq, name).First(&techscien).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

package repository

import (
	"errors"

	"gorm.io/gorm"
	"main.go/internal/models"
	"main.go/pkg/db"
)

func CreateSocialScien(socialscien *models.SocialScien) error {
	existing, err := socialScienExists(socialscien.Name)
	if err != nil {
		return err
	}
	if existing {
		return errors.New("socialscien exists")
	}
	sqlReq := `INSERT INTO socialscien(name) VALUES(?)`
	err = db.GetDB().Exec(sqlReq, socialscien.Name).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteSocialScienById(socialscienId int64) error {
	var socialscien *models.SocialScien
	if err := db.GetDB().First(&socialscien, socialscienId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	err := db.GetDB().Delete(&socialscien, socialscienId).Error
	if err != nil {
		return err
	}
	return nil
}

func GetSocialSciences() (socialsciences []*models.SocialScien, err error) {
	sqlReq := `SELECT * FROM socialscien ORDER BY name`
	err = db.GetDB().Raw(sqlReq).Scan(&socialsciences).Error
	if err != nil {
		return nil, err
	}
	return socialsciences, err
}

func GetSocialScienById(socialscienId int64) (socialscien *models.SocialScien, err error) {
	sqlReq := `SELECT * FROM socialscien WHERE id = ?`
	err = db.GetDB().Raw(sqlReq, socialscienId).Scan(&socialscien).Error
	if err != nil {
		return nil, err
	}
	return socialscien, nil
}

func socialScienExists(name string) (bool, error) {
	var socialscien *models.SocialScien
	sqlReq := `SELECT * FROM socialscien WHERE name = ?`
	err := db.GetDB().Raw(sqlReq, name).First(&socialscien).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

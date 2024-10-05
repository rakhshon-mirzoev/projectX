package repository

import (
	"errors"

	"gorm.io/gorm"
	"main.go/internal/models"
	"main.go/pkg/db"
)

func CreateSchool(school *models.School) (err error) {
	existing, err := schoolExists(school.Name)
	if err != nil {
		return err
	}
	if existing {
		return errors.New("school exists")
	}
	sqlReq := `INSERT INTO school(name, type, language, address, freelaunch, main_disc, transport, edonish, private, amount, shift) VALUES(?,?,?,?,?,?,?,?,?,?,?)`
	err = db.GetDB().Exec(sqlReq, school.Name, school.Type, school.Language, school.Address, school.FreeLaunch, school.MainDisc, school.Transport, school.EDonish, school.Private, school.Amount, school.Shift).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteSchoolById(schId int64) error {
	var school *models.School
	if err := db.GetDB().First(&school, schId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	err := db.GetDB().Delete(&school, schId).Error
	if err != nil {
		return err
	}
	return nil
}

func GetSchools(f *models.SchoolFilter) (schools *[]models.School, err error) {
	query := db.GetDB().Table("school s")
	if f.Name != nil {
		query.Where("s.name ilike ?", "%"+*f.Name+"%")
	}
	if f.Address != nil {
		query.Where("s.address ilike ?", "%"+*f.Address+"%")
	}
	if f.Language != nil {
		query.Where("s.language = ?", "%"+*f.Language+"%")
	}
	if f.Private != nil {
		query.Where("s.private = ?", f.Private)
	}
	if f.CityId != nil {
		query.Where("s.city_id = ?", f.CityId)
	}
	if f.EDonish != nil {
		query.Where("s.edonish = ?", f.EDonish)
	}
	if f.FreeLaunch != nil {
		query.Where("s.freelaunch = ?", f.FreeLaunch)
	}
	if f.MainDisc != nil {
		query.Where("s.main_disc ilike ?", f.MainDisc)
	}
	if f.Transport != nil {
		query.Where("s.transport = ?", f.Transport)
	}
	if f.Type != nil {
		query.Where("s.type ilike ?", f.Type)
	}
	err = query.Select("*").Order("name").Scan(&schools).Error
	if err != nil {
		return nil, err
	}
	return schools, nil
}

func GetSchoolById(schId int64) (school *models.School, err error) {
	sqlReq := `SELECT * FROM school WHERE id = ?`
	err = db.GetDB().Raw(sqlReq, schId).Scan(&school).Error
	if err != nil {
		return nil, err
	}
	return school, nil
}

func schoolExists(name string) (bool, error) {
	var school models.School
	sqlReq := `SELECT * FROM school WHERE name = ?`
	err := db.GetDB().Raw(sqlReq, name).First(&school).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

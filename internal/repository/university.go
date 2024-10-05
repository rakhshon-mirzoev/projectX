package repository

import (
	"errors"

	"gorm.io/gorm"
	"main.go/internal/models"
	"main.go/pkg/db"
)

var (
	ex     []int64
	exact  models.ExactScien
	social models.SocialScien
	tech   models.TechScien
	human  models.HumanScien
)

func CreateUni(uni *models.University) error {
	existingLogin, err := universityLoginExists(uni.Login)
	if err != nil {
		return err
	}
	if existingLogin {
		return errors.New("university login exists")
	}
	existing, err := universityExists(uni.Name)
	if err != nil {
		return err
	}
	if existing {
		return errors.New("university exists")
	}
	sqlReq := `INSERT INTO university(name,address,phone,email,amount,language,shift,system,private, login, password) VALUES(?,?,?,?,?,?,?,?,?,?,?)`
	err = db.GetDB().Exec(sqlReq, uni.Name, uni.Address, uni.Phone, uni.Email, uni.Amount, uni.Language, uni.Shift, uni.System, uni.Private, uni.Login, uni.Password).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUniById(uniId int64) error {
	var uni *models.University
	if err := db.GetDB().First(&uni, uniId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	err := db.GetDB().Delete(&uni, uniId).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUnis(f *models.UniversityFilter) (unis *[]models.University, err error) {
	query := db.GetDB().Table("university u")
	if f.Name != nil {
		query.Where("u.name ilike ?", "%"+*f.Name+"%")
	}
	if f.Address != nil {
		query.Where("u.address ilike ?", "%"+*f.Address+"%")
	}
	if f.Language != nil {
		query.Where("u.language = ?", "%"+*f.Language+"%")
	}
	if f.Private != nil {
		query.Where("u.private = ?", f.Private)
	}
	if f.CityId != nil {
		query.Where("u.city_id = ?", f.CityId)
	}
	err = query.Select("*").Order("name").Scan(&unis).Error
	if err != nil {
		return nil, err
	}
	return unis, nil
}

func GetUniById(uniId int64) (uni *models.University, err error) {
	err = db.GetDB().Where("id = ?", uniId).First(&uni).Error
	if err != nil {
		return nil, translateError(err)
	}
	return uni, nil
}

func GetUniExactSubs(uniId int64) (exsubs []models.ExactScien, err error) {
	sqlReq := `SELECT exactscien_id FROM universityexacts where university_id = ?`
	err = db.GetDB().Raw(sqlReq, uniId).Scan(&ex).Error
	if err != nil {
		return nil, err
	}
	for _, u := range ex {
		sqlReq = `SELECT name FROM exactscien WHERE id = ?`
		err = db.GetDB().Raw(sqlReq, u).Scan(&exact).Error
		if err != nil {
			return nil, err
		}
		exsubs = append(exsubs, exact)
	}
	return exsubs, nil
}

func GetUniHumanSubs(uniId int64) (husubs []models.HumanScien, err error) {
	sqlReq := `SELECT humanscien_id FROM universityhumans where university_id = ?`
	err = db.GetDB().Raw(sqlReq, uniId).Scan(&ex).Error
	if err != nil {
		return nil, err
	}
	for _, u := range ex {
		sqlReq = `SELECT name FROM humanscien WHERE id = ?`
		err = db.GetDB().Raw(sqlReq, u).Scan(&human).Error
		if err != nil {
			return nil, err
		}
		husubs = append(husubs, human)
	}
	return husubs, nil
}

func GetUniTechSubs(uniId int64) (tesubs []models.TechScien, err error) {
	sqlReq := `SELECT techscien_id FROM universitytechs where university_id = ?`
	err = db.GetDB().Raw(sqlReq, uniId).Scan(&ex).Error
	if err != nil {
		return nil, err
	}
	for _, u := range ex {
		sqlReq = `SELECT name FROM techscien WHERE id = ?`
		err = db.GetDB().Raw(sqlReq, u).Scan(&tech).Error
		if err != nil {
			return nil, err
		}
		tesubs = append(tesubs, tech)
	}
	return tesubs, nil
}

func GetUniSocialSubs(uniId int64) (sosubs []models.SocialScien, err error) {
	sqlReq := `SELECT socialscien_id FROM universitysocials where university_id = ?`
	err = db.GetDB().Raw(sqlReq, uniId).Scan(&ex).Error
	if err != nil {
		return nil, err
	}
	for _, u := range ex {
		sqlReq = `SELECT name FROM socialscien WHERE id = ?`
		err = db.GetDB().Raw(sqlReq, u).Scan(&social).Error
		if err != nil {
			return nil, err
		}
		sosubs = append(sosubs, social)
	}
	return sosubs, nil
}

func GetUniWithCityName() (universities []models.University, err error) {
	if err = db.GetDB().Preload("City").Find(&universities).Error; err != nil {
		return nil, err
	}
	return universities, nil
}

func ChangeUniversity(uid int64, f models.UniversityUpdateFilter) (uni *models.University, err error) {
	query := db.GetDB().Table("university u")
	if err := query.Where("u.id = ?", uid).First(&uni).Error; err != nil {
		return nil, err
	}
	if err := query.Model(&uni).Updates(f).Error; err != nil {
		return nil, err
	}
	return uni, nil
}

func universityExists(name string) (bool, error) {
	var uni models.University
	sqlReq := `SELECT * FROM university WHERE name = ?`
	err := db.GetDB().Raw(sqlReq, name).First(&uni).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func universityLoginExists(login string) (bool, error) {
	var uni models.University
	sqlReq := `SELECT * FROM university WHERE login = ?`
	err := db.GetDB().Raw(sqlReq, login).First(&uni).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

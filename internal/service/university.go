package service

import (
	log "main.go/logger"

	"main.go/internal/models"
	"main.go/internal/repository"
)

func CreateUni(uni *models.University) error {
	err := repository.CreateUni(uni)
	if err != nil {
		if err.Error() == "university exists" {
			log.Warn.Println("University with that name already exists!")
			return err
		} else if err.Error() == "record not found" {
			return nil
		} else {
			log.Error.Println("Error in repository (CreateUni): ", err)
			return err
		}
	} else {
		log.Info.Println("Successfully created a new university.")
	}
	return nil
}

func DeleteUniById(uniId int64) error {
	err := repository.DeleteUniById(uniId)
	if err != nil {
		log.Error.Println("Error in repository (DeleteUniById): ", err)
		return err
	}
	return nil
}

func GetUnis(f *models.UniversityFilter) (unis *[]models.University, err error) {
	unis, err = repository.GetUnis(f)
	if err != nil {
		log.Error.Println("Error in repository (GetUnis): ", err)
		return nil, err
	}
	return unis, nil
}

func GetUniById(uniId int64) (uni *models.University, err error) {
	uni, err = repository.GetUniById(uniId)
	if err != nil {
		log.Error.Println("Error in repository (GetUniById): ", err)
		return nil, err
	}
	return uni, nil
}

func GetUniExactSubs(uniId int64) (exsubs []models.ExactScien, err error) {
	exsubs, err = repository.GetUniExactSubs(uniId)
	if err != nil {
		log.Error.Println("Error in repository (GetUniExactSubs): ", err)
		return nil, err
	}
	return exsubs, nil
}

func GetUniHumanSubs(uniId int64) (husubs []models.HumanScien, err error) {
	husubs, err = repository.GetUniHumanSubs(uniId)
	if err != nil {
		log.Error.Println("Error in repository (GetUniHumanSubs): ", err)
		return nil, err
	}
	return husubs, nil
}

func GetUniTechSubs(uniId int64) (tesubs []models.TechScien, err error) {
	tesubs, err = repository.GetUniTechSubs(uniId)
	if err != nil {
		log.Error.Println("Error in repository (GetUniTechSubs): ", err)
		return nil, err
	}
	return tesubs, nil
}

func GetUniSocialSubs(uniId int64) (sosubs []models.SocialScien, err error) {
	sosubs, err = repository.GetUniSocialSubs(uniId)
	if err != nil {
		log.Error.Println("Error in repository (GetUniSocialSubs): ", err)
		return nil, err
	}
	return sosubs, nil
}

func GetUniWithCityName() (universitiesWithCityNames []models.UniversityWithCityName, err error) {
	universities, err := repository.GetUniWithCityName()
	if err != nil {
		log.Error.Println("Error in repository (GetUniWithCityName): ", err)
		return nil, err
	}
	for _, u := range universities {
		universitiesWithCityNames = append(universitiesWithCityNames, models.UniversityWithCityName{
			Id:   u.Id,
			Name: u.Name,
			City: u.City.Name,
		})
	}
	return universitiesWithCityNames, nil
}

func ChangeUniversity(uid int64, f models.UniversityUpdateFilter) (uni *models.University, err error) {
	uni, err = repository.ChangeUniversity(uid, f)
	if err != nil {
		log.Error.Println("Error in repository (ChangeUniversity): ", err)
		return nil, err
	}
	return uni, nil
}

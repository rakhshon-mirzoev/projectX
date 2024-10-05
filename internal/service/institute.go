package service

import (
	log "main.go/logger"

	"main.go/internal/models"
	"main.go/internal/repository"
)

func GetInstitutionList(t models.Institute, filU *models.UniversityFilter, filS *models.SchoolFilter) (sc *[]models.School, u *[]models.University, err error) {
	if t.Type == "school" {
		sc, _, err = repository.GetInstitutionList(t, filU, filS)
		if err != nil {
			log.Error.Println("Error in repository (GetInstitutionList): ", err)
			return nil, nil, err
		}
		return sc, u, nil
	} else if t.Type == "university" {
		_, u, err = repository.GetInstitutionList(t, filU, filS)
		if err != nil {
			log.Error.Println("Error in repository (GetInstitutionList): ", err)
			return nil, nil, err
		}
		return sc, u, nil
	} else if t.Type == "" {
		sc, u, err = repository.GetInstitutionList(t, filU, filS)
		if err != nil {
			log.Error.Println("Error in repository (GetInstitutionList): ", err)
			return nil, nil, err
		}
		return sc, u, nil
	} else {
		log.Error.Println("NIL DENIED")
		return nil, nil, err
	}
}

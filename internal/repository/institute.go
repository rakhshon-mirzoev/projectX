package repository

import "main.go/internal/models"

func GetInstitutionList(t models.Institute, filU *models.UniversityFilter, filS *models.SchoolFilter) (sc *[]models.School, u *[]models.University, err error) {
	if t.Type == "school" {
		sc, err = GetSchools(filS)
		if err != nil {
			return nil, nil, err
		}
		return sc, nil, nil
	} else if t.Type == "university" {
		u, err = GetUnis(filU)
		if err != nil {
			return nil, nil, err
		}
		return nil, u, nil
	} else if t.Type == "" {
		sc, err = GetSchools(filS)
		if err != nil {
			return nil, nil, err
		}
		u, err = GetUnis(filU)
		if err != nil {
			return nil, nil, err
		}
		return sc, u, nil
	} else {
		return nil, nil, err
	}
}

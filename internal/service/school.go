package service

import (
	log "main.go/logger"

	"main.go/internal/models"
	"main.go/internal/repository"
)

func CreateSchool(school *models.School) error {
	err := repository.CreateSchool(school)
	if err != nil {
		if err.Error() == "school exists" {
			log.Warn.Println("School with that name already exists!")
			return err
		} else if err.Error() == "record not found" {
			return nil
		} else {
			log.Error.Println("Error in repository (CreateSchool):", err)
			return err
		}
	} else {
		log.Info.Println("Successfully created a new chool.")
	}
	return nil
}

func DeleteSchoolById(schId int64) error {
	err := repository.DeleteSchoolById(schId)
	if err != nil {
		log.Error.Println("Error in repository (DeleteSchool):", err)
		return err
	}
	return nil
}

func GetSchools(f *models.SchoolFilter) (schools *[]models.School, err error) {
	schools, err = repository.GetSchools(f)
	if err != nil {
		log.Error.Println("Error in repository (GetSchools):", err)
		return nil, err
	}
	return schools, nil
}

func GetSchoolById(schId int64) (school *models.School, err error) {
	school, err = repository.GetSchoolById(schId)
	if err != nil {
		log.Error.Println("Error in repository (GetSchoolById):", err)
		return nil, err
	}
	return school, nil
}

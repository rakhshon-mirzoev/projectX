package models

type UniversityExacts struct {
	Id           int64      `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	UniversityId int64      `json:"university_id" gorm:"university_id"`
	University   University `json:"-" gorm:"foreignkey:UniversityId"`
	ExactScienId int64      `json:"exactscien_id" gorm:"column:exactscien_id"`
	ExactScien   ExactScien `json:"-" gorm:"foreignkey:ExactScienId"`
}

func (u *UniversityExacts) TableName() string {
	return "universityexacts"
}

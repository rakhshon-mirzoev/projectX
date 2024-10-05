package models

type UniversityHumans struct {
	Id           int64      `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	UniversityId int64      `json:"university_id" gorm:"university_id"`
	University   University `json:"-" gorm:"foreignkey:UniversityId"`
	HumanScienId int64      `json:"humanscien_id" gorm:"column:humanscien_id"`
	HumanScien   HumanScien `json:"-" gorm:"foreignkey:HumanScienId"`
}

func (u *UniversityHumans) TableName() string {
	return "universityhumans"
}

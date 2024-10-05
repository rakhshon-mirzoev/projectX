package models

type UniversityTechs struct {
	Id           int64      `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	UniversityId int64      `json:"university_id" gorm:"university_id"`
	University   University `json:"-" gorm:"foreignkey:UniversityId"`
	TechScienId  int64      `json:"techscien_id" gorm:"column:techscien_id"`
	TechScien    TechScien  `json:"-" gorm:"foreignkey:TechScienId"`
}

func (u *UniversityTechs) TableName() string {
	return "universitytechs"
}

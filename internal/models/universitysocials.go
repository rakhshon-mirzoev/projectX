package models

type UniversitySocials struct {
	Id            int64       `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	UniversityId  int64       `json:"university_id" gorm:"university_id"`
	University    University  `json:"-" gorm:"foreignkey:UniversityId"`
	SocialScienId int64       `json:"socialscien_id" gorm:"column:socialscien_id"`
	SocialScien   SocialScien `json:"-" gorm:"foreignkey:SocialScienId"`
}

func (u *UniversitySocials) TableName() string {
	return "universitysocials"
}

package models

type FavouriteUnis struct {
	Id           int64      `json:"id" gorm:"column:id"`
	UserId       int64      `json:"user_id" gorm:"user_id"`
	User         User       `json:"-" gorm:"foreignkey:UserId"`
	UniversityId int64      `json:"university_id" gorm:"university_id"`
	University   University `json:"-" gorm:"foreignkey:UniversityId"`
}

func (f *FavouriteUnis) TableName() string {
	return "favouriteunis"
}

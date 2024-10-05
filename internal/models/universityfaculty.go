package models

type UniversityFaculty struct {
	Id           int64      `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	UniversityId int64      `json:"university_id" gorm:"university_id"`
	University   University `json:"-" gorm:"foreignkey:UniversityId"`
	FacultyId    int64      `json:"faculty_id" gorm:"column:faculty_id"`
	Faculty      Faculty    `json:"-" gorm:"foreignkey:FacultyId"`
}

func (u *UniversityFaculty) TableName() string {
	return "universityfaculty"
}

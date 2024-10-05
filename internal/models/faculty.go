package models

type Faculty struct {
	Id   int64  `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Name string `json:"name" gorm:"column:name"`
}

func (f *Faculty) TableName() string {
	return "faculty"
}

package models

type TechScien struct {
	Id   int64  `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Name string `json:"name" gorm:"column:name"`
}

func (t *TechScien) TableName() string {
	return "techscien"
}

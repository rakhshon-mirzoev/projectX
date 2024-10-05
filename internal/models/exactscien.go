package models

type ExactScien struct {
	Id   int64  `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Name string `json:"name" gorm:"column:name"`
}

func (e *ExactScien) TableName() string {
	return "exactscien"
}

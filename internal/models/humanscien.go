package models

type HumanScien struct {
	Id   int64  `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Name string `json:"name" gorm:"column:name"`
}

func (s *HumanScien) TableName() string {
	return "humanscien"
}

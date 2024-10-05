package models

type Role struct {
	ID          int64  `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

func (r *Role) TableName() string {
	return "roles"
}

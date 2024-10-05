package models

type City struct {
	Id   int64  `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Name string `json:"name" gorm:"column:name"`
}

func (c *City) TableName() string {
	return "city"
}

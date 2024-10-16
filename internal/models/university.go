package models

type University struct {
	Id             int64  `json:"id" gorm:"column:id;autoIncrement;primary_key"`
	Description    string `json:"description" gorm:"column:description"`
	Name           string `json:"name" gorm:"column:name"`
	Address        string `json:"address" gorm:"address"`
	Phone          string `json:"phone" gorm:"column:phone"`
	Email          string `json:"email" gorm:"column:email"`
	Amount         int64  `json:"amount" gorm:"amount"`
	Language       string `json:"language" gorm:"column:language"`
	Shift          string `json:"shift" gorm:"column:shift"`
	System         string `json:"system" gorm:"column:system"`
	Private        bool   `json:"private" gorm:"column:private"`
	RoleId         int64  `json:"role_id" gorm:"column:role_id"`
	Role           Role   `json:"-" gorm:"foreignkey:RoleId"`
	CityId         int64  `json:"city_id" gorm:"column:city_id"`
	City           City   `json:"-" gorm:"foreignkey:CityId"`
	Login          string `json:"login" gorm:"column:login"`
	Password       string `json:"password" gorm:"column:password"`
	Identification bool   `json:"identification" gorm:"column:identification"`
}

func (u *University) TableName() string {
	return "university"
}

type UniversityFilter struct {
	Name     *string `json:"name" form:"name"`
	Address  *string `json:"address" form:"address"`
	Language *string `json:"language" form:"language"`
	Private  *bool   `json:"private" form:"private"`
	CityId   *int    `json:"city_id" form:"city_id"`
}

type UniversityUpdateFilter struct {
	Phone       *string `json:"phone" form:"phone"`
	Amount      *int64  `json:"amount" form:"amount"`
	Language    *string `json:"language" form:"language"`
	System      *string `json:"system" form:"system"`
	Shift       *string `json:"shift" form:"shift"`
	Description *string `json:"description" form:"description"`
}

type UniversityWithCityName struct {
	Id   int64  `json:""`
	Name string `json:"name"`
	City string `json:"city"`
}

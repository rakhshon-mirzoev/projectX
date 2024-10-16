package models

import "time"

type User struct {
	Id        int64      `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Name      string     `json:"name" gorm:"column:name"`
	Phone     string     `json:"phone" gorm:"column:phone"`
	BirthDate *time.Time `json:"birthDate" gorm:"column:birthdate"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Age       int        `json:"age" gorm:"column:age"`
	Login     string     `json:"login" gorm:"column:login"`
	Password  string     `json:"password" gorm:"column:password"`
	RoleId    int64      `json:"role" gorm:"column:role_id"`
	Role      Role       `json:"-" gorm:"foreignkey:RoleId"`
	CityName  string     `json:"city" gorm:"-"`
	CityId    int64      `json:"-" gorm:"column:city_id"`
	City      City       `json:"-" gorm:"foreignkey:CityId"`
	UserType  string     `json:"user_type" gorm:"column:user_type"`
	Active    bool       `json:"active" gorm:"column:active"`
}

func (u *User) TableName() string {
	return "users"
}

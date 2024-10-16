package models

type School struct {
	Id             int64  `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Name           string `json:"name" gorm:"column:name"`
	Type           string `json:"type" gorm:"column:type"`
	Language       string `json:"language" gorm:"column:language"`
	Address        string `json:"address" gorm:"column:address"`
	FreeLaunch     bool   `json:"freeLaunch" gorm:"column:freelaunch"`
	Transport      bool   `json:"transport" grom:"column:transport"`
	EDonish        bool   `json:"eDonish" gorm:"column:edonish"`
	Private        bool   `json:"private" gorm:"column:private"`
	Amount         int64  `json:"amount" gorm:"column:amount"`
	Shift          string `json:"shift" gorm:"column:shift"`
	Phone          string `json:"phone" gorm:"column:phone"`
	Email          string `json:"email" gorm:"column:email"`
	CityId         int64  `json:"city_id" gorm:"city_id"`
	City           City   `json:"-" gorm:"foreignkey:CityId"`
	RoleId         int64  `json:"role_id" gorm:"column:role_id"`
	Role           Role   `json:"-" gorm:"foreignkey:RoleId"`
	Login          string `json:"login" gorm:"column:login"`
	Password       string `json:"password" gorm:"column:password"`
	Identification bool   `json:"identification" gorm:"column:identification"`
}

func (s *School) TableName() string {
	return "school"
}

type SchoolFilter struct {
	Name       *string `json:"name" form:"name"`
	Type       *string `json:"type" form:"type"`
	Language   *string `json:"language" form:"language"`
	Address    *string `json:"address" form:"language"`
	MainDisc   *string `json:"main_disc" form:"main_disc"`
	Private    *bool   `json:"private" form:"private"`
	CityId     *int    `json:"city_id" form:"city_id"`
	EDonish    *bool   `json:"edonish" form:"edonish"`
	Transport  *bool   `json:"transport" form:"transport"`
	FreeLaunch *bool   `json:"freelaunch" form:"freelaunch"`
}

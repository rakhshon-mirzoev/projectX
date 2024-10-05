package models

type Institute struct {
	Type string `json:"type" form:"type" default:"university"`
}

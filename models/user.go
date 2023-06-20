package models

import "gorm.io/gorm"

type Calendar struct {
	gorm.Model
	Id           int    `json:"ID" gorm:"primary_key"`
	Day          string `json:"day"`
	Month        string `json:"month"`
	Year         string `json:"year"`
	Numberofdays int    `json:"int"`
}

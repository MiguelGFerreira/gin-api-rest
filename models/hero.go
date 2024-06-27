package models

import "gorm.io/gorm"

type Hero struct {
	gorm.Model
	Name  string `json:"name"`
	Ssn   string `json:"ssn"`
	Power string `json:"power"`
}

var Heroes []Hero

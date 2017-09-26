package models

import "github.com/jinzhu/gorm"

type Consumption struct {
	gorm.Model
	Liter string `json:"liter,omitempty"`
	Month string `json:"month,omitempty"`
	Year  string `json:"year,omitempty"`
}

type ConsumptionList []Consumption

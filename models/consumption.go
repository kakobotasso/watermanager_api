package models

type Consumption struct {
	Id    int    `json:"id,omitempty"`
	Liter string `json:"liter,omitempty"`
	Month string `json:"month,omitempty"`
}

type ConsumptionList []Consumption

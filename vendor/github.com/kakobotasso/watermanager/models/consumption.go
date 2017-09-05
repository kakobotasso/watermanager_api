package models

type Consumption struct {
	Id    int    `json:"id,omitempty"`
	Liter string `json:"liter,omitempty"`
}

type ConsumptionList []Consumption

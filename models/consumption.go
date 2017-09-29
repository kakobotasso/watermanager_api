package models

type Consumption struct {
	Model
	Liter  string `json:"liter,omitempty"`
	Month  string `json:"month,omitempty"`
	Year   string `json:"year,omitempty"`
	Serial string `json:"serial,omitempty"`
}

type ConsumptionList []Consumption

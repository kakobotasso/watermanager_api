package models

type Error struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

type Errors []Error

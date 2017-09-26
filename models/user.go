package models

type User struct {
	Model
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Cpf      string `json:"cpf,omitempty"`
	Username string `gorm:"unique" json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
}

type Users []User

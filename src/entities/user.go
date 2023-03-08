package entities

import "github.com/Jacobbrewer1/chess-boards/src/custom"

type User struct {
	Id        int             `json:"id,omitempty"`
	FirstName string          `json:"firstName,omitempty"`
	Surname   string          `json:"surname,omitempty"`
	Email     string          `json:"email,omitempty"`
	Password  string          `json:"password,omitempty"`
	LastLogin custom.Datetime `json:"lastLogin,omitempty"`
}

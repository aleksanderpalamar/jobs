package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role        string
	Level       string
	Company     string
	Location    string
	Remote      bool
	Link        string
	Salary      string
	Title       string
	Description string
	Keywords    string
}

package schemas

import (
	"time"

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
	Salary      int64
	Title       string
	Description string
	Keywords    string
}

type OpeningResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	Role        string    `json:"role"`
	Level       string    `json:"level"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	Remote      bool      `json:"remote"`
	Link        string    `json:"link"`
	Salary      int64     `json:"salary"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Keywords    string    `json:"keywords"`
}

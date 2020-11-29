package models

import (
	"time"

	"gorm.io/gorm"
)

// Citizen Model
type Citizen struct {
	gorm.Model
	FullName   string
	DateBirth  time.Time `json:"date_birth"`
	NoKk       string
	NoKtp      string
	PlaceBirth string
	Address    string
	AccountID  uint
}

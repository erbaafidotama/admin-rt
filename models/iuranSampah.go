package models

import (
	"time"

	"gorm.io/gorm"
)

// IuranSampah Model
type IuranSampah struct {
	gorm.Model
	AccountID   uint
	PayDate     time.Time `json:"pay_date"`
	Description string
}

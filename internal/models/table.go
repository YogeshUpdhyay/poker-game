package models

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	Name     string `json:"name"`
	MaxBuyIn int    `json:"maxBuyIn"`
}

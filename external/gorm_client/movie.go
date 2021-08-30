package gorm_client

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string
	Slug        string `gorm:"unique"`
	Description string
	Duration    int
	Image       string
}

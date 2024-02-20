package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Weight int `json:"weight"`
	Height int `json:"height"`
}

package models

import "github.com/jinzhu/gorm"

//Artist model
type Artist struct {
	gorm.Model
	Nombre string  `json:"nombre"`
	Image  string  `json:"image,omitempty"`
	Albums []Album `json:"albums,omitempty"`
}

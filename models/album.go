package models

import "github.com/jinzhu/gorm"

//Album model
type Album struct {
	gorm.Model
	Nombre   string `json:"nombre"`
	Year     uint16 `json:"year"`
	Image    string `json:"image,omitempty"`
	ArtistID uint   `json:"artistId"`
	Artist   Artist `json:"artist,omitempty"`
}

package models

import "github.com/jinzhu/gorm"

type Location struct {
	gorm.Model
	GeoX float64 `json:"GeoX"`
	GeoY float64 `json:"GeoY"`
	Name string  `json:"Name"`
}

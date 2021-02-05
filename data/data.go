package data

import (
	"postgis_test/models"

	"github.com/jinzhu/gorm"
)

func CreateTestData(db *gorm.DB) {
	data := []models.Location{
		{
			GeoX: 51.990830,
			GeoY: 5.967340,
			Name: "Looierstraat 5",
		},
		{
			GeoX: 51.990830,
			GeoY: 5.966880,
			Name: "Looierstraat X",
		},
		{
			GeoX: 51.994330,
			GeoY: 5.981940,
			Name: "Kastanjehof Velp",
		},
		{
			GeoX: 52.087921,
			GeoY: 6.157130,
			Name: "Somewhere else",
		},
	}

	for _, d := range data {
		db.Create(&d)
	}
}

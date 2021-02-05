package main

import (
	"fmt"
	"os"
	"postgis_test/data"
	"postgis_test/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func main() {
	// initalize gorm db
	x := Initalize()
	if x != nil {
		fmt.Println(x)
	}
	data.CreateTestData(db)

	var list []models.Location
	// create query strings
	q := ComposeQuery(51.990830, 5.967340, 5000)
	qPag := ComposeQueryLimit(51.990830, 5.967340, 5000, 10, 0)

	db.Raw(q).Scan(&list)
	db.Raw(qPag).Scan(&list)

	for _, l := range list {
		fmt.Println(l)
	}
	os.Exit(0)
}

func ComposeQuery(x, y float64, radius int) string {
	return fmt.Sprintf(`SELECT * FROM locations
	WHERE ST_DWithin(ST_MakePoint(geo_x,geo_y)::geography,ST_MakePoint(%f, %f)::geography,%d)
	LIMIT 1000
	`, x, y, radius)
}
func ComposeQueryLimit(x, y float64, radius, limit, offset int) string {
	return fmt.Sprintf(`SELECT * FROM locations
	WHERE ST_DWithin(ST_MakePoint(geo_x,geo_y)::geography,ST_MakePoint(%f, %f)::geography,%d)
	LIMIT %d
	OFFSET %d
	`, x, y, radius, limit, offset)
}

func Initalize() error {
	username := "test"
	password := "test"
	dbName := "test"
	dbHost := "localhost"

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Println("We cant op open a DATABASE")
		return err
	}

	db = conn
	db = db.Debug()

	db.AutoMigrate(&models.Location{})

	return nil
}

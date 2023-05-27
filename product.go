package main

import (
	"errors"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type Product struct {
	Name     string
	Category string
	SKU      string
}

var db *gorm.DB
var err error

var (
	exampleProducts = []Product{
		{Name: "Widget", Category: "Widgets", SKU: "WID001"},
		{Name: "Thingy", Category: "Thingies", SKU: "TNG001"},
		{Name: "Gadget", Category: "Gadgets", SKU: "GDT001"},
	}
)

func initDB() *gorm.DB {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=product sslmode=disable")

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	return db
}

func SaveTestData() {
	for index := range exampleProducts {
		db.Create(&exampleProducts[index])
	}
}

func (p *Product) save() error {
	existingProducts := get("SKU", p.SKU)
	if len(existingProducts) != 0 {
		log.Errorf("Product already exists: (duplicate SKU: %s)", p.SKU)
		return errors.New("product already exists (duplicate SKU)")
	}
	db.Create(&p)
	return nil
}

func searchAll(searchTerm string) []Product {
	var results []Product
	db.Where("name LIKE ? OR category LIKE ? OR sku LIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%", "%"+searchTerm+"%").Find(&results)
	return results
}

func search(fieldName string, fieldValue string) []Product {
	var results []Product
	db.Where(fieldName+" LIKE ?", "%"+fieldValue+"%").Find(&results)
	return results
}

func get(fieldName string, fieldValue string) []Product {
	var products []Product
	db.Where(fieldName+" = ?", fieldValue).Find(&products)
	return products
}

func getAll() []Product {
	var products []Product
	db.Find(&products)
	return products
}

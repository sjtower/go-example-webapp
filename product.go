package main

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Product struct {
	Name     string
	Category string
	SKU      string
}

// todo: map by SKU for quicker lookups
var products []Product

func (p *Product) save() error {
	if get(p.SKU) != nil {
		log.Errorf("Product already exists: (duplicate SKU: %s)", p.SKU)
		return errors.New("product already exists (duplicate SKU)")
	}
	products = append(products, *p)
	return nil
}

func searchAll(searchTerm string) []Product {
	var results []Product
	for _, p := range products {
		if strings.Contains(p.SKU, searchTerm) || strings.Contains(p.Name, searchTerm) || strings.Contains(p.Category, searchTerm) {
			results = append(results, p)
		}
	}
	return results
}

func searchName(searchTerm string) []Product {
	var results []Product
	for _, p := range products {
		if strings.Contains(p.Name, searchTerm) {
			results = append(results, p)
		}
	}
	return results
}

func searchCategory(searchTerm string) []Product {
	var results []Product
	for _, p := range products {
		if strings.Contains(p.Category, searchTerm) {
			results = append(results, p)
		}
	}
	return results
}

func searchSKU(searchTerm string) []Product {
	var results []Product
	for _, p := range products {
		if strings.Contains(p.SKU, searchTerm) {
			results = append(results, p)
		}
	}
	return results
}

func get(sku string) *Product {
	for _, p := range products {
		if p.SKU == sku {
			return &p
		}
	}
	return nil
}

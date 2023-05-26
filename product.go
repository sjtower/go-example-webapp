package main

import (
	"errors"
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
		return errors.New("product already exists (duplicate SKU)")
	}
	products = append(products, *p)
	return nil
}

func search(searchTerm string) []Product {
	var results []Product
	for _, p := range products {
		if strings.Contains(p.SKU, searchTerm) || strings.Contains(p.Name, searchTerm) || strings.Contains(p.Category, searchTerm) {
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

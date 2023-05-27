package main

import "testing"

func TestSearchAll_empty(t *testing.T) {
	p2 := searchAll("testSKU0")
	if len(p2) != 0 {
		t.Errorf("expected no results, got %v", len(p2))
	}
}

func TestSearchAll(t *testing.T) {
	saveTestProduct(t, "TestProductAll")
	p2 := searchAll("testSKU0")
	if p2[0].SKU != "testSKU0TestProductAll" {
		t.Errorf("Expected SKU: %s, got: %s", "testSKU0", p2[0].SKU)
	}
}

func TestSearchName(t *testing.T) {
	saveTestProduct(t, "TestProductName")
	p2 := searchName("TestProduct")
	if p2[0].Name != "TestProductAll" {
		t.Errorf("Expected Name: %s, got: %s", "TestProduct", p2[0].Name)
	}
}

func TestSearchCategory(t *testing.T) {
	saveTestProduct(t, "TestProductCat")
	p2 := searchCategory("This is a sample Product.")
	if p2[0].Category != "This is a sample Product." {
		t.Errorf("Expected Category: %s, got: %s", "This is a sample Product.", p2[0].Category)
	}
}

func TestSearchSKU(t *testing.T) {
	saveTestProduct(t, "TestProductSKU")
	p2 := searchSKU("testSKU0")
	if p2[0].SKU != "testSKU0TestProductAll" {
		t.Errorf("Expected SKU: %s, got: %s", "testSKU0", p2[0].SKU)
	}
}

func saveTestProduct(t *testing.T, name string) {
	sku := "testSKU0" + name
	p1 := &Product{Name: name, Category: "This is a sample Product.", SKU: sku}
	err := p1.save()
	if err != nil {
		t.Errorf("Error saving product: %v", err)
	}
}

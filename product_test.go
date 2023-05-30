package main

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestAddAndSearch(t *testing.T) {
	saveTestProduct(t, "Widget", "WID001")

	//make request to search for the product
	response, err := http.PostForm("http://localhost:8080/search/", url.Values{"name": {"Widget"}})

	if err != nil {
		t.Errorf("Error making search request: %v", err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	if !strings.Contains(string(body), "<div>Name: Widget</div><div>Category: This is a sample Product.</div><div>SKU: WID001</div>") {
		t.Errorf("Expected Widget result, got %s", string(body))
	}
	//fmt.Printf("%s\n", string(body))
}

func saveTestProduct(t *testing.T, name string, sku string) {
	response, err := http.PostForm("http://localhost:8080/save/", url.Values{"name": {"Widget"},
		"category": {"This is a sample Product."}, "sku": {"WID001"}})
	if err != nil {
		t.Errorf("Error making add request: %v", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil || body == nil {
		t.Errorf("Error reading response body: %v", err)
	}
}

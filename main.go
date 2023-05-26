package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func test() {
	p1 := &Product{Name: "TestProduct", Category: "This is a sample Product.", SKU: "testSKU0"}
	err := p1.save()
	if err != nil {
		log.Errorf("Error saving product: %v", err)
	}
	p2 := get("testSKU0")
	log.Infof("Product: %v", p2.SKU)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	//drop the leading "/search/" part of the URL path to get the search term
	searchTerm := r.URL.Path[len("/search/"):]
	results := search(searchTerm)
	for _, p := range results {
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div><div>%s</div>", p.Name, p.Category, p.SKU)
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	test()
	http.HandleFunc("/", handler)
	http.HandleFunc("/search/", searchHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

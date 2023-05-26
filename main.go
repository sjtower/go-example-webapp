package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"html/template"
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

func addProductHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "add")
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	category := r.FormValue("category")
	sku := r.FormValue("sku")
	log.Infof("New Product saved - name: %s, category: %s, sku: %s", name, category, sku)
	p := &Product{Name: name, Category: category, SKU: sku}
	err := p.save()
	if err != nil {
		//409 response code indicates duplicate record
		//todo: error handling
		http.Redirect(w, r, "/add", 409)
	}
	http.Redirect(w, r, "/add", http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, nil)
}

func main() {
	test()
	http.HandleFunc("/", handler)
	http.HandleFunc("/search/", searchHandler)
	http.HandleFunc("/add/", addProductHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

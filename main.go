package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("add.html", "search.html"))

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

	name := r.FormValue("name")
	category := r.FormValue("category")
	sku := r.FormValue("sku")
	allFieldsTerm := r.FormValue("searchAll")
	log.Infof("Search - name: %s, category: %s, sku: %s, searchAll: %s", name, category, sku, allFieldsTerm)

	var results []Product
	if allFieldsTerm != "" {
		results = searchAll(allFieldsTerm)
	} else {
		if name != "" {
			results = searchName(name)
		} else if category != "" {
			results = searchCategory(category)
		} else if sku != "" {
			results = searchSKU(sku)
		}
	}

	for _, p := range results {
		fmt.Fprintf(w, "<div>Name: %s</div><div>Category: %s</div><div>SKU: %s</div>", p.Name, p.Category, p.SKU)
	}

	renderTemplate(w, "search")
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/add", http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	test()
	http.HandleFunc("/", handler)
	http.HandleFunc("/search/", searchHandler)
	http.HandleFunc("/add/", addProductHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

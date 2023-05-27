package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("add.html", "search.html"))

func baseHandler(w http.ResponseWriter, r *http.Request) {
	var products = getAll()
	err := json.NewEncoder(w).Encode(&products)
	if err != nil {
		log.Errorf("Error encoding exampleProducts: %s", err.Error())
	}
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
			results = search("Name", name)
		} else if category != "" {
			results = search("Category", category)
		} else if sku != "" {
			results = search("SKU", sku)
		}
	}

	//todo: replace this with an html template
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

	db := initDB()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Errorf("Error closing db: %s", err.Error())
		}
	}(db)

	http.HandleFunc("/", baseHandler)
	http.HandleFunc("/search/", searchHandler)
	http.HandleFunc("/add/", addProductHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/test/", testHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func testHandler(writer http.ResponseWriter, request *http.Request) {
	SaveTestData()
}

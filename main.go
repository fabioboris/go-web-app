package main

import (
	"html/template"
	"net/http"

	product "github.com/fabioboris/go-web-app/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := product.GetAllProducts()
	templates.ExecuteTemplate(w, "Index", products)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

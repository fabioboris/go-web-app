package products

import (
	"html/template"
	"net/http"

	product "github.com/fabioboris/go-web-app/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := product.GetAllProducts()
	templates.ExecuteTemplate(w, "Index", products)
}

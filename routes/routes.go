package routes

import (
	"net/http"

	products "github.com/fabioboris/go-web-app/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", products.Index)
}

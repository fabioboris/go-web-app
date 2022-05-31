package main

import (
	"net/http"

	"github.com/fabioboris/go-web-app/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}

package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func connectToDB() *sql.DB {
	connStr := "user=fabioboris dbname=playground sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	CheckError(err)
	return db
}

type Product struct {
	Name  string
	Price float64
	Qty   int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := connectToDB()
	defer db.Close()

	rows, err := db.Query("SELECT name, price, quantity FROM products")
	CheckError(err)

	products := []Product{}

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.Name, &p.Price, &p.Qty)
		CheckError(err)
		products = append(products, p)
	}

	templates.ExecuteTemplate(w, "Index", products)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

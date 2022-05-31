package product

import "github.com/fabioboris/go-web-app/db"

type Product struct {
	Name  string
	Price float64
	Qty   int
}

func GetAllProducts() []Product {
	db := db.ConnectToDB()
	defer db.Close()

	rows, err := db.Query("SELECT name, price, quantity FROM products")
	if err != nil {
		panic(err)
	}

	products := []Product{}
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.Name, &p.Price, &p.Qty)
		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}
	return products
}

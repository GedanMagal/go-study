package main

import (
	"crud-app-go/models/products"
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func dbConnection() *sql.DB {
	connection := "user=postgres dbname=alura_loja password=online host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// products := []products.Product{
	// 	{Name: "Camiseta Vans", Description: "Camiseta preta", Price: 100., Quantity: 100},
	// 	{Name: "Camiseta Adidas", Description: "Camiseta preta", Price: 150.1, Quantity: 100},
	// }
	db := dbConnection()

	productsQuery, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := products.Product{}
	products := []products.Product{}

	for productsQuery.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productsQuery.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)

	}

	templates.ExecuteTemplate(w, "Index", products)

	defer db.Close()
}

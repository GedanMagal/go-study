package models

import (
	"crud-app-go/db"
	"log"
)

type Product struct {
	id                int
	Name, Description string
	Price             float64
	Quantity          int
}

func GetAllProducts() []Product {

	// products := []products.Product{
	// 	{Name: "Camiseta Vans", Description: "Camiseta preta", Price: 100., Quantity: 100},
	// 	{Name: "Camiseta Adidas", Description: "Camiseta preta", Price: 150.1, Quantity: 100},
	// }

	db := db.Connection()

	productsQuery, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

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

	defer db.Close()

	return products
}

func InsertProduct(name, description string, price float64, quantity int) {

	db := db.Connection()

	insertStatement, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")

	if err != nil {
		log.Println(err.Error())
	}

	insertStatement.Exec(name, description, price, quantity)

	defer db.Close()
}

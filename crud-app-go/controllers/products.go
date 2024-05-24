package controllers

import (
	models "crud-app-go/models/products"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	products := models.GetAllProducts()

	templates.ExecuteTemplate(w, "Index", products)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {

	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		priceConverted, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println(err.Error())
		}

		quantityConverted, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println(err.Error())
		}

		models.InsertProduct(name, description, priceConverted, quantityConverted)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

package routes

import (
	"crud-app-go/controllers"
	"net/http"
)

func LoadRoutes() {

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.NewProduct)
	http.HandleFunc("/insert", controllers.Insert)
}

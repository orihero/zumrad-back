package router

import (
	"zumrad/controller"
	"zumrad/env"

	"github.com/gorilla/mux"
)

func NewRouter() {
	env.Router = mux.NewRouter()
	env.Router.HandleFunc("/products", controller.GetProducts).Methods("GET")
	env.Router.HandleFunc("/products", controller.CreateProduct).Methods("POST")
	env.Router.HandleFunc("/order", controller.CreateOrder).Methods("POST")
	env.Router.HandleFunc("/order", controller.GetOrders).Methods("GET")
}

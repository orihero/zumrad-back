package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zumrad/database"
	"zumrad/models"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	var orders []models.Order
	connection.Preload("Client").Preload("Location").Find(&orders)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		fmt.Println(err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}

	connection.Create(&order)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)

}

package controller

import (
	"encoding/json"

	"net/http"
	"zumrad/database"
	"zumrad/models"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	var products []models.Product
	connection.Preload("Color").Preload("Characteristic").Preload("Sizes").Find(&products)
	w.Header().Set("Content-Type", "applicaation/json")
	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}

	connection.Create(&product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}

package routes

import (
	"encoding/json"
	"net/http"
	"stock_information_project/db"
	"stock_information_project/models"
)

func GetStocksHandler(w http.ResponseWriter, r *http.Request) {
	var stocks []models.Stock
	db.DB.Find(&stocks)
	json.NewEncoder(w).Encode(stocks)

}

func CreateStockHandler(w http.ResponseWriter, r *http.Request) {
	var stocks models.Stock
	json.NewDecoder(r.Body).Decode(&stocks)
	createdStock := db.DB.Create(&stocks)
	err := createdStock.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&stocks)
}

func GetStockHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteStockHandler(w http.ResponseWriter, r *http.Request) {

}

package routes

import (
	"encoding/json"
	"net/http"
	"stock_information_project/db"
	"stock_information_project/models"

	"github.com/gorilla/mux"
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
	var stock models.Stock
	params := mux.Vars(r)

	db.DB.First(&stock, params["id"])

	if stock.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Stock not found"))
		return
	}
	json.NewEncoder(w).Encode(stock)

}

func DeleteStockHandler(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	params := mux.Vars(r)

	db.DB.Delete(&stock, params["id"])
	w.WriteHeader(http.StatusNoContent)

	if stock.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Stock not found"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

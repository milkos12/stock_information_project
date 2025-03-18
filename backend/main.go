package main

import (
	"net/http"
	"stock_information_project/db"
	"stock_information_project/models"
	routes "stock_information_project/routers"

	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Stock{})

	r := mux.NewRouter()

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//Stocks ROuters
	r.HandleFunc("/stocks", routes.GetStocksHandler).Methods("GET")
	r.HandleFunc("/stocks", routes.CreateStockHandler).Methods("POST")
	r.HandleFunc("/stocks/{id}", routes.GetStockHandler).Methods("GET")
	r.HandleFunc("/stocks/{id}", routes.DeleteStockHandler).Methods("DELETE")

	//Fetch
	r.HandleFunc("/fetch", routes.GetListStocksHandler).Methods("GET")

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}

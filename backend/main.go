package main

import (
	"net/http"
	"stock_information_project/db"
	"stock_information_project/models"
	routes "stock_information_project/routers"

	"log"

	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	log.Println("🚀 Servidor iniciando...")

	db.DBConnection()

	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Stock{})

	r := mux.NewRouter()
	r.Use(enableCORS)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//Stocks ROuters
	r.HandleFunc("/stocks", routes.GetStocksHandler).Methods("GET")
	r.HandleFunc("/stocks", routes.CreateStockHandler).Methods("POST")
	r.HandleFunc("/stocks/{id}", routes.GetStockHandler).Methods("GET")
	r.HandleFunc("/stocks/{id}", routes.DeleteStockHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}

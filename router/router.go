package router

import (
	"github.com/gorilla/mux"
	"github.com/v4r5v4m5/go-stocks-api/middleware"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET", "OPTIONS")       // get stock details
	r.HandleFunc("/api/stocks", middleware.GetAllStock).Methods("GET", "OPTIONS")        // get all stock details
	r.HandleFunc("/api/newstock", middleware.CreateStock).Methods("POST", "OPTIONS")     // create new stock
	r.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")    // edit a stock
	r.HandleFunc("/api/stock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS") // delete a stock

	return r
}

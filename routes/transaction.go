package routes

import (
	"dewetour/handlers"
	// "dewetour/pkg/middleware"
	"dewetour/pkg/middleware"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gorilla/mux"
)

func TransacRoute(r *mux.Router) {
	transacRepository := repositories.RepositoryTransac(mysql.DB)
	h := handlers.HandleTransac(transacRepository)

	r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
	// r.HandleFunc("/trip/{id}", h.UpdateTrip).Methods("PATCH")
	// r.HandleFunc("/trip/{id}", h.DeleteTrip).Methods("DELETE")
	// r.HandleFunc("/trips", h.FindTrip).Methods("GET")
	// r.HandleFunc("/trip/{id}", h.GetTrip).Methods("GET")
}

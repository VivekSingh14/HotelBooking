package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Community/hotel-booking/config"
	"github.com/gorilla/mux"
)

func RegisterHandlers(router *mux.Router, config *config.Config) {

	router.HandleFunc("/health", HealthCheck).Methods(http.MethodGet)
	http.Handle("/", router)

	bookingRouter := router.PathPrefix("/booking/v1").Subrouter()

	bookingEndpoints(bookingRouter)
}

func bookingEndpoints(bookingRouter *mux.Router) {
	bookingRouter.HandleFunc("/health", HealthCheck).Methods(http.MethodGet)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

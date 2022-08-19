package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Community/hotel-booking/config"
	"github.com/gorilla/mux"
)

func RegisterHandlers(router *mux.Router, config *config.Config) {
	service := NewBookingController(context.Background(), config)

	bookingRouter := router.PathPrefix("/booking/v1").Subrouter()

	bookingEndpoints(bookingRouter, service)
}

func bookingEndpoints(bookingRouter *mux.Router, service *BookingService) {
	bookingRouter.HandleFunc("/health", check).Methods(http.MethodGet)
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health Check </h1>")
}

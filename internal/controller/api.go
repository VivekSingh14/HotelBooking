package controller

import (
	"net/http"

	"github.com/Community/hotel-booking/config"
	"github.com/gorilla/mux"
)

func RegisterHandlers(router *mux.Router, config *config.Config) {

	service := NewBookingService(config)

	bookingRouter := router.PathPrefix("/booking/v1").Subrouter()

	bookingEndpoints(bookingRouter, service)
}

func bookingEndpoints(bookingRouter *mux.Router, service *BookingService) {
	bookingRouter.HandleFunc("/health", service.HealthCheck).Methods(http.MethodGet)
}

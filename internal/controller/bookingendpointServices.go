package controller

import (
	"fmt"
	"log"
	"net/http"
)

func (s *BookingService) HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

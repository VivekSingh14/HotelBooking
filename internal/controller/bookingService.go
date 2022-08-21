package controller

import "github.com/Community/hotel-booking/config"

type BookingService struct {
	Config *config.Config
}

func NewBookingService(config *config.Config) *BookingService {
	service := BookingService{Config: config}
	return &service
}

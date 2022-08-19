package controller

import (
	"context"

	"github.com/Community/hotel-booking/config"
)

type BookingService struct {
	Config *config.Config
}

type BookingDAO struct {
}

func NewBookingController(ctx context.Context, config *config.Config) *BookingService {
	service := BookingService{Config: config}
	return &service
}

package controller

import (
	"github.com/Community/hotel-booking/config"
)

type BookingService struct {
	Config *config.Config
}

// type DAOService struct {
// 	BookingRepository dao.BooKingRepository
// }

func NewBookingService(config *config.Config) *BookingService {
	service := BookingService{Config: config}
	return &service
}

// func (b *BookingService) GetDAOService() DAOService {

// 	dao := DAOService{
// 		BookingRepository: b.,
// 	}

// 	return dao

// }

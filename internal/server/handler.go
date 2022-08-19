package server

import "github.com/Community/hotel-booking/internal/controller"

func (s *Server) registerHandlers() {
	controller.RegisterHandlers(s.router, s.config)
}

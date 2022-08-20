package server

import "github.com/Community/hotel-booking/internal/controller"

func (s *Server) registerHandler() {
	controller.RegisterHandlers(s.router, s.config)
}

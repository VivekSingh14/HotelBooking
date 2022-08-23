package controller

import (
	"fmt"
	"net/http"
)

func (s *BookingService) HealthCheck(w http.ResponseWriter, r *http.Request) {

	//daoService := s.GetDAOService()

	//context := gorillactx.Get(r, "temp").(context.Context)
	//context := context.Background().Done()

	// res, err := daoService.BookingRepository.GetHealthCheck()

	// if err != nil {
	// 	log.Printf("error occured %d", err)
	// 	return
	// }

	//log.Printf("entering health check end point %v \n", res)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Health is good.")
}

package main

import (
	"fmt"
	"log"

	"github.com/Community/hotel-booking/config"
	"github.com/Community/hotel-booking/internal/server"
)

func main() {

	log.Println("Starting api server.")

	cfg, err := config.GetConfig()

	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	} else {
		log.Println("config loaded successfully.")
	}
	fmt.Printf("AppVersion: %s \n", cfg.Server.AppVersion)

	srv := server.NewServer(cfg)

	srv.ListenAndServe()

}

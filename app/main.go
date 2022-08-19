package main

import (
	"fmt"
	"log"

	"github.com/Community/hotel-booking/config"
	"github.com/Community/hotel-booking/helper/signals"
	"github.com/Community/hotel-booking/internal/server"
)

//adding comment
// func index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "<h1>Hello World </h1>")
// }

// func check(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "<h1>Health Check </h1>")
// }

func main() {

	log.Println("Starting api server.")

	cfg, err := config.GetConfig()

	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	} else {
		fmt.Println("config loaded successfully.")
	}
	fmt.Printf("AppVersion: %s \n", cfg.Server.AppVersion)
	s := server.NewServer(cfg)

	stopch := signals.SetupHandler()

	s.ListenAndServe(stopch)

	// http.HandleFunc("/", index)
	// http.HandleFunc("/health", check)
	// fmt.Println("Server starting...")
	// http.ListenAndServe(":5000", nil)
}

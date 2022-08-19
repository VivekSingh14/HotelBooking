package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Community/hotel-booking/config"
)

//adding comment
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World </h1>")
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health Check </h1>")
}

func main() {

	log.Println("Starting api server.")

	cfg, err := config.GetConfig()

	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	} else {
		fmt.Println("config loaded successfully.")
	}
	fmt.Printf("AppVersion: %s \n", cfg.Server.AppVersion)
	http.HandleFunc("/", index)
	http.HandleFunc("/health", check)
	fmt.Println("Server starting...")
	http.ListenAndServe(":5000", nil)
}

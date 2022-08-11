package main

import (
	"fmt"
	"net/http"
)

//adding comment
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World </h1>")
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health Check </h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/health", check)
	fmt.Println("Server starting...")
	http.ListenAndServe(":5000", nil)
}

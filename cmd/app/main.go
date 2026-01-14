package main

import (
	"log"
	"net/http"

	"simple-cicd/cmd/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.HealthCheck)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
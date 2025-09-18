package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"jenkins-cicd-demo/internal/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Setup routes
	http.HandleFunc("/", handlers.SimpleHelloHandler)
	http.HandleFunc("/api/hello", handlers.HelloHandler)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/dashboard", handlers.Dashboard)
	http.HandleFunc("/metrics", handlers.Metrics)

	fmt.Printf("Server starting on port %s\n", port)
	fmt.Println("Endpoints:")
	fmt.Println("   GET /          - Simple hello message")
	fmt.Println("   GET /api/hello - JSON hello response")
	fmt.Println("   GET /health    - Health check")
	fmt.Println("   GET /dashboard - CI/CD Pipeline Dashboard")
	fmt.Println("   GET /metrics   - Pipeline metrics")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
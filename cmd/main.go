package main

import (
	"fmt"
	"log"
	"net/http"
	"retail-pulse-image-processor/internal/handlers"
	"retail-pulse-image-processor/internal/services"
)

func main() {
	err := services.LoadStoreMaster("config/StoreMasterAssignment.csv")
	if err != nil {
		log.Fatalf("Error loading StoreMaster: %v\n", err)
	}

	http.HandleFunc("/api/submit", handlers.SubmitJobHandler)
	http.HandleFunc("/api/status", handlers.JobStatusHandler)

	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

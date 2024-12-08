package main

import (
	"log"
	"net/http"

	"receipt-processor/internal/handlers"
)

func main() {
	http.HandleFunc("/receipts/process", handlers.ProcessReceipt)
	http.HandleFunc("/receipts/{id}/points", handlers.GetPoints)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

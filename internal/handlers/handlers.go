package handlers

import (
	"encoding/json"
	"net/http"

	"receipt-processor/internal/services"

	"../models"

	"github.com/google/uuid"
)

var receiptData = make(map[string]models.Receipt)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	receiptData[id] = receipt

	response := models.ProcessResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	receipt, exists := receiptData[id]

	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	points := services.CalculatePoints(receipt)
	response := models.PointsResponse{Points: points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

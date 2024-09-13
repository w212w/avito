package handlers

import (
	"encoding/json"
	"net/http"
	"tender-service/models"
	"tender-service/repository"
)

func NewBidHandler(w http.ResponseWriter, r *http.Request) {
	var newBid models.Bid
	if err := json.NewDecoder(r.Body).Decode(&newBid); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := repository.CreateBid(&newBid); err != nil {
		http.Error(w, "Ошибка при создании предложения", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBid)
}

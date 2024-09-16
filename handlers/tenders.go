package handlers

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"tender-service/models"
	"tender-service/repository"
)

func NewTenderHandler(w http.ResponseWriter, r *http.Request) {
	var newTender models.Tender
	if err := json.NewDecoder(r.Body).Decode(&newTender); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := repository.CreateTender(&newTender); err != nil {
		http.Error(w, "Ошибка при создании тендера", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTender)
}

func TenderListHandler(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	serviceTypes := r.URL.Query()["service_type"]

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	tenders, err := repository.GetTenders(limit, offset, serviceTypes)
	if err != nil {
		http.Error(w, "Ошибка при получении списка тендеров", http.StatusInternalServerError)
		return
	}

	sort.Sort(models.ByName(tenders))

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tenders)
}

func UpdateTenderStatusHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		TenderID string `json:"tenderId"`
		Status   string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := repository.UpdateTenderStatus(requestData.TenderID, requestData.Status); err != nil {
		http.Error(w, "Ошибка при обновлении статуса тендера", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(requestData)
}

func GetUserTendersHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	limitParam := r.URL.Query().Get("limit")
	offsetParam := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetParam)
	if err != nil || offset < 0 {
		offset = 0
	}

	tenders, err := repository.GetTendersByUsername(username, limit, offset)
	if err != nil {
		http.Error(w, "Failed to fetch tenders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tenders)
}

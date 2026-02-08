package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
	"time"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) HandleTransactions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetTransactions(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetReport(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.service.GetTransactionToday()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func (h *TransactionHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	_, err1 := time.Parse("2006-01-02", startDate)
	_, err2 := time.Parse("2006-01-02", endDate)
	if err1 != nil || err2 != nil {
		http.Error(w, "format tanggal harus YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	transactions, err := h.service.GetReport(startDate, endDate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req models.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	transaction, err := h.service.Checkout(req.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

func (h *TransactionHandler) SalesSummaryToday(w http.ResponseWriter, r *http.Request) {

	transaction, err := h.service.SalesSummaryToday()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

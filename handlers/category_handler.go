package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
	"strconv"
	"strings"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// HandleCategories - GET, POST /api/categories/
func (h *CategoryHandler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.GetCategories(w, r)
	} else if r.Method == http.MethodPost {
		h.CreateCategory(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetCategories - GET /api/categories/ -> menampilkan semua data kategori
func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	Categories, err := h.service.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Categories)
}

// CreateCategory - POST /api/categories/ -> menambah kategori baru
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var Category models.Category
	err := json.NewDecoder(r.Body).Decode(&Category)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.service.CreateCategory(&Category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Category)
}

// HandleCategoryById - GET,PUT,DELETE /api/categories/{id}
func (h *CategoryHandler) HandleCategoryByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.GetCategoryByID(w, r)
	} else if r.Method == http.MethodPut {
		h.UpdateCategoryByID(w, r)
	} else if r.Method == http.MethodDelete {
		h.DeleteCategoryByID(w, r)
	}
}

// GetCategoryByID - GET /api/categories/{id}
func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}
	Category, err := h.service.GetCategoryByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Category)
}

// UpdateCategoryByID - PUT /api/categories/{id}
func (h *CategoryHandler) UpdateCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var Category models.Category
	err = json.NewDecoder(r.Body).Decode(&Category)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	Category.ID = id
	err = h.service.UpdateCategoryByID(&Category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Category)
}

// DeleteCategoryByID - DELETE /api/categories/{id}
func (h *CategoryHandler) DeleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteCategoryByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Category deleted successfully",
	})
}

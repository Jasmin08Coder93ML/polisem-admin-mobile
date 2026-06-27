package products

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler { return &Handler{service: s} }

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	json.NewDecoder(r.Body).Decode(&p)
	
	err := h.service.CreateProduct(p)
	if err != nil {
		http.Error(w, "Error saving", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

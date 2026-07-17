package warehouse

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) UpdateStockHandler(w http.ResponseWriter, r *http.Request) {
	// Логика парсинга JSON и вызова h.service.AdjustStock
}

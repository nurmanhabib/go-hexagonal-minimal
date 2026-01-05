package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"hexagonal-minimal/internal/domain/user"
)

type Handler struct {
	service *user.Service
}

func NewHandler(s *user.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	err := h.service.Create(r.Context(), req.Name)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)

	u, err := h.service.Get(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	json.NewEncoder(w).Encode(u)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)

	err := h.service.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

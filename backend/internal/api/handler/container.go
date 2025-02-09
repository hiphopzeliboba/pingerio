package handler

import (
	"encoding/json"
	"net/http"
	"pingerio/backend/internal/model"
	"pingerio/backend/internal/service"
)

type ContainerHandler struct {
	service service.ContainerService
}

func NewContainerHandler(service service.ContainerService) *ContainerHandler {
	return &ContainerHandler{
		service: service,
	}
}

func (h *ContainerHandler) GetContainers(w http.ResponseWriter, r *http.Request) {
	containers, err := h.service.GetContainers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(containers)
}

func (h *ContainerHandler) SaveContainers(w http.ResponseWriter, r *http.Request) {
	var containers []model.Container
	if err := json.NewDecoder(r.Body).Decode(&containers); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.SaveContainers(r.Context(), containers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

package server

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

type health struct {
	Status bool `json:"status"`
}

type healthHandler struct {
	db *gorm.DB
}

func (h *healthHandler) Live(w http.ResponseWriter, r *http.Request) {
	h.sendResponse(w, true)
}

func (h *healthHandler) Ready(w http.ResponseWriter, r *http.Request) {
	db, err := h.db.DB()
	if err != nil {
		h.sendResponse(w, false)
		return
	}

	err = db.Ping()
	if err != nil {
		h.sendResponse(w, false)
		return
	}

	h.sendResponse(w, true)
}

func (h *healthHandler) sendResponse(w http.ResponseWriter, status bool) {
	resp := health{Status: status}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

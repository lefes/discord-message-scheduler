package http

import (
	"encoding/json"
	"net/http"

	"github.com/lefes/discord-message-scheduler/pkg/logger"
)

func (s *Server) liveHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "ok"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Error(err)
		return
	}
}

// TODO: add services to check if the server is ready
func (s *Server) readyHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "ok"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Error(err)
		return
	}
}

func (s *Server) RegisterHandlers() {
	http.HandleFunc("/live", s.liveHandler)
	http.HandleFunc("/ready", s.readyHandler)
}

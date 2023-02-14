package api

import (
	"encoding/json"
	"net/http"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Version(w http.ResponseWriter, r *http.Request) {
	ver := map[string]string{
		"version": "1.0.0",
	}
	jsonContent, err := json.Marshal(ver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonContent)
}

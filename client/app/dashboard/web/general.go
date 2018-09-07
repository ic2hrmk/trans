package web

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeResponse(data interface{}, code int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	responseBody, err := json.MarshalIndent(data, "   ", "  ")
	if err != nil {
		log.Println("[web-dashboard] failed to marshal response")
		return
	}

	if _, err = w.Write(responseBody); err != nil {
		log.Println("[web-dashboard] failed to write response")
		return
	}
}

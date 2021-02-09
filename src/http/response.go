package http

import (
	"encoding/json"
	"net/http"
)

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *customResponseWriter {
	return &customResponseWriter{w, http.StatusOK}
}

func (crw *customResponseWriter) WriteHeader(code int) {
	crw.statusCode = code
	crw.ResponseWriter.WriteHeader(code)
}

type errorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// NewErrorResponse - HTTP Response handling for errors
func NewErrorResponse(w http.ResponseWriter, statusCode int, response string) {
	error := errorResponse{
		true,
		response,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&error)
	return
}

// ResponseJSON - HTTP Success Response
func ResponseJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)
}

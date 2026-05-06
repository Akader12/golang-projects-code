package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool   `json:"successs"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func SendSuccesss(w http.ResponseWriter,message string, data any) {
	w.Header().Set("Content/Type","application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: message,
		Data: data,
	})
}

func SendError(w http.ResponseWriter, message string, err string, code int)  {
	w.Header().Set("Content/Type","application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(APIResponse{
		Success: false,
		Message: message,
		Error: err,
	})
}
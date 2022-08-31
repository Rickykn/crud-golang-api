package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type jsonResponse struct {
	Code    int
	Message string
	Data    interface{}
}

func WriteResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := &jsonResponse{
		Code:    statusCode,
		Message: message,
		Data:    data,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}
}

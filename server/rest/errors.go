package rest

import (
	"encoding/json"
	"net/http"
)

type restError struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

func NewBadRequest(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	res := restError{
		StatusCode: http.StatusBadRequest,
		Message:    err.Error(),
	}

	json.NewEncoder(w).Encode(res)
}

func NewSuccessful(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

package services

import (
	"encoding/json"
	"net/http"
)

func GetJob(w http.ResponseWriter, r *http.Request) {
	var response = struct {
		Message string `json:"message"`
	}{
		Message: "Getting Job",
	}

	message, err := json.Marshal(response)

	if err != nil {
		panic("Error marshalling ErrorResponse")
	}

	w.Write(message)
}

func PostJob(w http.ResponseWriter, r *http.Request) {
	var response = struct {
		Message string `json:"message"`
	}{
		Message: "Posting Job",
	}

	message, err := json.Marshal(response)

	if err != nil {
		panic("Error marshalling ErrorResponse")
	}

	w.Write(message)
}

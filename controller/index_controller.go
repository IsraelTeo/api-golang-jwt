package controller

import (
	"encoding/json"
	"net/http"
)

type Gretting struct {
	Message    string `json:"mesage"`
	StatusCode int    `json: code`
}

func InitRoute(w http.ResponseWriter, r *http.Request) {
	greeting := Gretting{
		Message:    "Hello world",
		StatusCode: 200,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(greeting)
}

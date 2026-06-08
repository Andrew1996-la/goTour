package http

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	task := Task{
		ID:    1,
		Title: "Learn Go",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(task)
}
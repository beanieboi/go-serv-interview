package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
)

type Response struct {
	ID string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Header().Set("Content-Type", "application/json")
		if rand.Intn(10) <= 1 {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		json.NewEncoder(w).Encode(Response{ID: id})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	http.ListenAndServe(":"+port, mux)
}

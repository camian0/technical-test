package main

import (
	"encoding/json"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Routes() *chi.Mux {
	mux := chi.NewMux()
	//middleware
	mux.Use(
		middleware.Logger,
		middleware.Recoverer,
	)
	mux.Get("/", helloHandler)
	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := map[string]interface{}{"message": "hello"}
	_ = json.NewEncoder(w).Encode(res)
}

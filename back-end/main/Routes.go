package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"io"
	"myapp.com/enron/helpers"
	"myapp.com/enron/model"
	"net/http"
)

const ZINCURL = "http://localhost:4080/api/enron/_search"

func Routes() *chi.Mux {
	mux := chi.NewMux()
	//middleware
	mux.Use(
		middleware.Logger,
		middleware.Recoverer,
	)
	//cors
	mux.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	mux.Get("/", helloHandler)
	mux.Post("/search", searchZicHandler)
	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := map[string]interface{}{"message": "hello"}
	_ = json.NewEncoder(w).Encode(res)
}

func searchZicHandler(res http.ResponseWriter, req *http.Request) {
	userEncoded := helpers.EncodeUser()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Error al leer el cuerpo de la solicitud", http.StatusInternalServerError)
		return
	}

	var bodyBytes = helpers.ConvertBytes(body)
	reqPost, err := http.NewRequest("POST", ZINCURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		http.Error(res, "Error al leer el cuerpo de la solicitud", http.StatusInternalServerError)
		return
	}
	reqPost.Header.Set("Content-Type", "application/json")
	reqPost.Header.Set("Authorization", "Basic "+userEncoded)

	client := &http.Client{}
	response, err := client.Do(reqPost)
	if err != nil {
		panic(err)
	}

	body, err = io.ReadAll(response.Body)
	if err != nil {
		http.Error(res, "Error al leer el cuerpo de la solicitud", http.StatusInternalServerError)
		return
	}

	var decode model.Response
	err = json.Unmarshal(body, &decode)
	if err != nil {
		fmt.Println("Error al decodificar JSON:", err)
		return
	}
	decodeData := map[string]interface{}{"data": decode}
	_ = json.NewEncoder(res).Encode(decodeData)
	defer response.Body.Close()
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
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
	res.Header().Set("Content-Type", "application/json")

	userEncoded := helpers.EncodeUser()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Error al leer el cuerpo de la solicitud", http.StatusInternalServerError)
		return
	}
	var jsonStr = []byte(string(body))
	reqPost, err := http.NewRequest("POST", ZINCURL, bytes.NewBuffer(jsonStr))
	reqPost.Header.Set("Content-Type", "application/json")
	reqPost.Header.Set("Authorization", "Basic "+userEncoded)
	if err != nil {
		http.Error(res, "Error al leer el cuerpo de la solicitud", http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	response, err := client.Do(reqPost)
	if err != nil {
		panic(err)
	}

	body, err = ioutil.ReadAll(response.Body)
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

	defer response.Body.Close()
	_ = json.NewEncoder(res).Encode(decodeData)
}

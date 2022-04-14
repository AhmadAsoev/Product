package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"product/cmd/handleFunc"
	"product/pkg/application/repository"
)

const (
	GET  = "GET"
	POST = "POST"
)

func main() {
	log.Println("Server connected")
	if err := repository.Connect(); err != nil {
		log.Fatal("Server can't connect to DB")
	}

	router := mux.NewRouter()
	port := ":8080"

	//Health check server
	router.HandleFunc("/health", handleFunc.Health).Methods(GET)

	//AddProduct
	router.HandleFunc("/product", handleFunc.AddProduct).Methods(POST)

	log.Println("server start!")
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("Server is not ready!")
	}
}

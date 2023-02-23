package main

import (
	"github.com/converge/simple-go-api/api"
	"log"
	"net/http"
)

func main() {

	// endpoints
	http.HandleFunc("/", api.Version)
	http.HandleFunc("/healthz", api.Healthz)

	service := http.Server{Addr: ":7001"}

	log.Println("API listening...")
	log.Panic(service.ListenAndServe(), nil)

}

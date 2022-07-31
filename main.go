package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("main endpoint :)")
		w.Write([]byte("hi!"))
	})
	service := http.Server{
		Addr: ":7001",
	}

	log.Println("API listening...")
	log.Panic(service.ListenAndServe(), nil)

}

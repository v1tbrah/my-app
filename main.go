package main

import (
	"log"
	"net/http"
	"time"
)

const readHeaderTimeout = time.Second * 5

func main() {
	log.Print("Hello, i'm started\n")

	server := http.Server{Addr: ":3333", ReadHeaderTimeout: readHeaderTimeout}

	handler := http.NewServeMux()

	handler.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusOK)
			response := []byte("Hello from API")
			if _, err := w.Write(response); err != nil {
				log.Printf("writing response: %s", err.Error())
			}
		},
	)

	server.Handler = handler

	log.Print("I'm getting started listen sever on :3333")

	if err := server.ListenAndServe(); err != nil {
		log.Printf("listening server :3333: %s", err.Error())
	}
}

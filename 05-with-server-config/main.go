package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	sr := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Millisecond,
		WriteTimeout:   10 * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(sr.ListenAndServe())
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func JsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Powered-By", "GoLang")
		next.ServeHTTP(w, r)
	})
}

func JsonResponseFunc(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{}
	response["message"] = "Hello World"
	json.NewEncoder(w).Encode(response)
}

func main() {
	handler := http.HandlerFunc(JsonResponseFunc)
	http.Handle("/", JsonMiddleware(handler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

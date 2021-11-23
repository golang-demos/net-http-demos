package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func JsonResponseFunc(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{}
	response["message"] = "Hello World"
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", JsonResponseFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

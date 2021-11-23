package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type CustomResponse struct {
	Data    []string
	IsError bool
}

func SendCustomResponse(w http.ResponseWriter, data []string) {
	response := CustomResponse{}
	response.Data = data
	json.NewEncoder(w).Encode(response)
}

func JsonResponseFunc(w http.ResponseWriter, r *http.Request) {
	messages := []string{"Hello World"}
	SendCustomResponse(w, messages)
}

func main() {
	http.HandleFunc("/", JsonResponseFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

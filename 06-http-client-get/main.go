package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal("Unable to fetch data")
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// String
	fmt.Println(string(body))

	// Struct
	var jsonBody Post
	json.NewDecoder(resp.Body).Decode(&jsonBody)
	fmt.Println("UserId : ", jsonBody.UserId)
	fmt.Println("Id : ", jsonBody.Id)
	fmt.Println("Title : ", jsonBody.Title)
	fmt.Println("Body : ", jsonBody.Body)
}

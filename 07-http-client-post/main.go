package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	resp, err := http.PostForm("https://jsonplaceholder.typicode.com/posts", url.Values{
		"userId": {"1"},
		"id":     {"1"},
		"title":  {"My dummy Post"},
		"body":   {"My dummy Post body"},
	})
	if err != nil {
		log.Fatal("Unable to Post data")
	}

	defer resp.Body.Close()

	// Struct
	var jsonBody Post
	json.NewDecoder(resp.Body).Decode(&jsonBody)
	fmt.Println("UserId : ", jsonBody.UserId)
	fmt.Println("Id : ", jsonBody.Id)
	fmt.Println("Title : ", jsonBody.Title)
	fmt.Println("Body : ", jsonBody.Body)
}

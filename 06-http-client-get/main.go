package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

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
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal("Unable to fetch data")
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}

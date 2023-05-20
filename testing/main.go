package main

import (
	"net/http"

	"github.com/aleroxac/learning-go/testing/api"
)

func main() {
	http.HandleFunc("/hello-world", api.HelloWorld)
	http.HandleFunc("/hello", api.HelloYou)

	http.ListenAndServe(":8080", nil)
}

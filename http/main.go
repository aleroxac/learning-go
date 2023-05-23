package main

import (
	"fmt"
	"io"
	"net/http"
)

func GetRequest() {
	req, err := http.Get("https://api.gameofthronesquotes.xyz/v1/random")
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", res)
}

func main() {
	GetRequest()
}

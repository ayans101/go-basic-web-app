package main

import (
	"fmt"
	"net/http"

	"github.com/ayans101/go-basic-web-app/pkg/handlers"
)

const portNumber = ":4000"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application on port", portNumber)
	http.ListenAndServe(portNumber, nil)

}

//	go run cmd/web/*.go

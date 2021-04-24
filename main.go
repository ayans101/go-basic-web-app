package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":4000"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)

}

//	go run *.go

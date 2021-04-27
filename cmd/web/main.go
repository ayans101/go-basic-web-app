package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ayans101/go-basic-web-app/pkg/config"
	"github.com/ayans101/go-basic-web-app/pkg/handlers"
	"github.com/ayans101/go-basic-web-app/pkg/render"
)

const portNumber = ":4000"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Println("Starting application on port", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}

//	go run cmd/web/*.go

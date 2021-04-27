package main

import (
	"net/http"

	"github.com/bmizerany/pat"

	"github.com/ayans101/go-basic-web-app/pkg/config"
	"github.com/ayans101/go-basic-web-app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}

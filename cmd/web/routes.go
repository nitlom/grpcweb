package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/nitlom/webserver/pkg/config"
	"github.com/nitlom/webserver/pkg/handlers"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux

}

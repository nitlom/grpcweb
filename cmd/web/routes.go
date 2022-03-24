package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nitlom/webserver/pkg/config"
	"github.com/nitlom/webserver/pkg/handlers"
)

func Routes(app *config.AppConfig) http.Handler {

	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux

}

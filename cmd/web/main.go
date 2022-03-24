package main

import (
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/nitlom/webserver/pkg/config"
	"github.com/nitlom/webserver/pkg/handlers"
	"github.com/nitlom/webserver/pkg/render"
)

var (
	httpAddress string
)

const portNumber = ":8080"

func main() {

	var listenAddress = ""
	if strings.HasPrefix(portNumber, ":") {
		listenAddress = "*" + portNumber
	}
	log.Println("Listening at", listenAddress)

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&app),
	}
	srv.ListenAndServe()

}

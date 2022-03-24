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

	httpAddress = ":8080"

	app.TemplateCache = tc
	render.NewTemplates(&app)

	serveHttp()
}

func serveHttp() {
	r := http.NewServeMux()
	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/about", handlers.About)

	httpServer := &http.Server{
		Handler: r,
		Addr:    portNumber,
	}

	//log.Printf("Starting web-server on %s", portNumber)
	log.Fatal(httpServer.ListenAndServe())
}

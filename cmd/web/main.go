package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/oz1tmm/webserver/pkg/config"
	"github.com/oz1tmm/webserver/pkg/handlers"
	"github.com/oz1tmm/webserver/pkg/render"
)

var (
	httpAddress string
)

const portNumber = ":8080"

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	serviceAddress := ":7000"
	_, err := net.Listen("tcp", serviceAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Listening at:", serviceAddress)

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

	log.Println("And here...")
	fmt.Printf("Starting web-server on %s", portNumber)
	log.Fatal(httpServer.ListenAndServe())
}

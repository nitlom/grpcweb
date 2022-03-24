package handlers

import (
	"net/http"

	"github.com/oz1tmm/webserver/pkg/render"
)

// Home is the handler for the Home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// Home is the handler for the Home page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

package controller

import (
	"html/template"
	"log"
	"net/http"
)

type errorHandler struct {
}

func (e *errorHandler) Create(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/error.html")
	if err != nil {
		log.Printf("could not read html: %v", err)
	}
	tmpl.Execute(w, "")
}
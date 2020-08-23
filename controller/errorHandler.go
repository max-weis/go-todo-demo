package controller

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type errorHandler struct {
}

func (e *errorHandler) Error(w http.ResponseWriter, r *http.Request) {
	var req = struct {
		Error []error
	}{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return
	}

	tmpl, err := template.ParseFiles("static/error.html")
	if err != nil {
		log.Printf("could not read html: %v", err)
	}
	tmpl.Execute(w, "")
}

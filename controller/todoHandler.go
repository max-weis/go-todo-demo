package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	gotodo "todo"
)

type todoHandler struct {
	s gotodo.Service
}

func (t *todoHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("Create")
}

func (t *todoHandler) FindById(w http.ResponseWriter, r *http.Request) {
	log.Printf("FindById")
}

func (t *todoHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	pLimit := params["limit"]
	pOffset := params["offset"]

	limit, err := strconv.Atoi(pLimit[0])
	if err != nil {

	}

	offset, err := strconv.Atoi(pOffset[0])
	if err != nil {

	}

	todos, err := t.s.FindAll(limit, offset)
	if err != nil {

	}

	var resp = struct {
		Todos []gotodo.Todo `json:"todos"`
	}{Todos: todos}

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Printf("could not read html: %v", err)
	}
	tmpl.Execute(w, resp)
}

func (t *todoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Printf("Delete")
}

func (t *todoHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Printf("Update")
}

package controller

import (
	"fmt"
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
	err := r.ParseForm()
	if err != nil {
		log.Printf("error occured: %+v", err)
		http.Redirect(w, r, "/error", 400)
		return
	}

	title := r.Form["title"][0]
	description := r.Form["description"][0]

	todo, err := t.s.Create(title, description)
	if err != nil {
		log.Printf("error occured: %+v", err)
		http.Redirect(w, r, "/error", 400)
		return
	}

	redirectURL := fmt.Sprintf("/todo/%d", todo.ID)

	http.Redirect(w, r, redirectURL, 301)
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

		Offset int
		Limit  int
	}{Todos: todos, Offset: offset, Limit: limit}

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

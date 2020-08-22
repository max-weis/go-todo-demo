package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"html/template"
	"log"
	"net/http"
	"strconv"
	gotodo "todo"
)

type todoHandler struct {
	s      gotodo.Service
	logger zap.Logger
}

func (t *todoHandler) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		t.logger.Info("could not parse form", zap.Error(err))
		return
	}

	title := r.Form["title"][0]
	description := r.Form["description"][0]

	todo, err := t.s.Create(title, description)
	if err != nil {
		t.logger.Info("could not create todo", zap.Error(err))
		return
	}

	redirectURL := fmt.Sprintf("/todo/%d", todo.ID)

	http.Redirect(w, r, redirectURL, 301)
}

func (t *todoHandler) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		t.logger.Info("could not read id", zap.Error(err))
		return
	}

	todo, err := t.s.FindById(id)
	if err != nil {
		t.logger.Info("could not find todo", zap.Error(err))
		return
	}

	var resp = struct {
		Todo gotodo.Todo
	}{Todo: todo}

	tmpl, err := template.ParseFiles("static/detail.html")
	if err != nil {
		t.logger.Info("could not parse html", zap.Error(err))
		return
	}

	tmpl.Execute(w, resp)
}

func (t *todoHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	pLimit := params["limit"]
	pOffset := params["offset"]

	limit, err := strconv.Atoi(pLimit[0])
	if err != nil {
		t.logger.Info("could not read limit", zap.Error(err))
		return
	}

	offset, err := strconv.Atoi(pOffset[0])
	if err != nil {
		t.logger.Info("could not read offset", zap.Error(err))
		return
	}

	todos, err := t.s.FindAll(limit, offset)
	if err != nil {
		t.logger.Info("could not find todos", zap.Error(err))
		return
	}

	var resp = struct {
		Todos []gotodo.Todo `json:"todos"`

		Offset int
		Limit  int
	}{Todos: todos, Offset: offset, Limit: limit}

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		t.logger.Info("could not parse file", zap.Error(err))
		return
	}

	tmpl.Execute(w, resp)
}

func (t *todoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		t.logger.Info("could not read id", zap.Error(err))
		return
	}

	d := json.NewDecoder(r.Body)

	var req = struct {
		Offset int
		Limit  int
	}{}

	err = d.Decode(&req)
	if err != nil {
		t.logger.Info("could not decode request", zap.Error(err))
		return
	}

	_, err = t.s.Delete(id)
	if err != nil {
		t.logger.Info("could not delete todo", zap.Error(err))
		return
	}
}

func (t *todoHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Printf("Update")
}

package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"html/template"
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
	id, ok := t.getID(r)
	if !ok {
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
	id, ok := t.getID(r)
	if !ok {
		return
	}

	decoder := json.NewDecoder(r.Body)

	var req = struct {
		Offset int
		Limit  int
	}{}

	err := decoder.Decode(&req)
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
	id, ok := t.getID(r)
	if !ok {
		return
	}

	decoder := json.NewDecoder(r.Body)

	var req = struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      bool   `json:"status"`
	}{}

	err := decoder.Decode(&req)
	if err != nil {
		t.logger.Info("could not decode request", zap.Error(err))
		return
	}

	_, err = t.s.Update(id, req.Title, req.Description, req.Status)
	if err != nil {
		t.logger.Info("could not update todo", zap.Error(err))
		return
	}
}

func (t *todoHandler) Done(w http.ResponseWriter, r *http.Request) {
	id, ok := t.getID(r)
	if !ok {
		return
	}

	params := r.URL.Query()["status"]

	status, err := strconv.ParseBool(params[0])
	if err != nil {
		t.logger.Info("could not read status", zap.Error(err))
		return
	}

	_, err = t.s.Done(id, status)
	if err != nil {
		t.logger.Info("could not set status of todo", zap.Error(err))
		return
	}
}

func (t *todoHandler) getID(r *http.Request) (int, bool) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		t.logger.Info("could not read id", zap.Error(err))
		return -1, false
	}

	return id, true
}

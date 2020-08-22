package controller

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"todo/todo"
)

type Server struct {
	Todo todo.Service

	logger zap.Logger
	Router *mux.Router
}

func NewServer(todoService todo.Service, logger zap.Logger, router *mux.Router) *Server {
	s := &Server{Todo: todoService, Router: router}
	h := todoHandler{s: todoService, logger: logger}
	e := errorHandler{}

	// redirect to list view
	s.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/todo?offset=0&limit=5", 301)
	}).Methods("GET")

	s.Router.HandleFunc("/todo/{id}", h.FindById).Methods("GET")
	s.Router.HandleFunc("/todo", h.FindAll).Methods("GET")
	s.Router.HandleFunc("/todo", h.Create).Methods("POST")
	s.Router.HandleFunc("/todo/{id}", h.Update).Methods("PUT")
	s.Router.HandleFunc("/todo/{id}", h.Done).Methods("PATCH")
	s.Router.HandleFunc("/todo/{id}", h.Delete).Methods("DELETE")

	s.Router.HandleFunc("/error", e.Error).Methods("GET")

	s.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

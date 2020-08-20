package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"todo/todo"
)

type Server struct {
	Todo todo.Service

	Router *mux.Router
}

func NewServer(todoService todo.Service, router *mux.Router) *Server {
	s := &Server{Todo: todoService, Router: router}
	h := todoHandler{s: todoService}

	s.Router.HandleFunc("/todo/{id}", h.FindById).Methods("GET")
	s.Router.HandleFunc("/todo", h.FindAll).Methods("GET")
	s.Router.HandleFunc("/todo", h.Create).Methods("POST")
	s.Router.HandleFunc("/todo/{id}", h.Update).Methods("PUT")
	s.Router.HandleFunc("/todo", h.Delete).Methods("DELETE")

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

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

	s.Router.HandleFunc("/", redirect).Methods("GET")

	s.Router.HandleFunc("/todo/{id}", h.FindById).Methods("GET")
	s.Router.HandleFunc("/todo", h.FindAll).Methods("GET", "DELETE") // redirects dont change method
	s.Router.HandleFunc("/todo", h.Create).Methods("POST")
	s.Router.HandleFunc("/todo/{id}", h.Update).Methods("PUT")
	s.Router.HandleFunc("/todo/{id}", h.Delete).Methods("DELETE")

	s.Router.HandleFunc("/error", e.Error).Methods("GET")

	staticDir := "/stylelib/"
	s.Router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("./static"+staticDir))))

	return s
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo?offset=0&limit=5", 301)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

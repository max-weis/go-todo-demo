package server

import (
	"github.com/gorilla/mux"
	"github.com/max-weis/go-todo-demo/todo"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

type Server struct {
	Todo todo.Service

	logger zap.Logger
	Router *mux.Router
	DB     *gorm.DB
}

func NewServer(todoService todo.Service, logger zap.Logger, router *mux.Router, db *gorm.DB) *Server {
	s := &Server{Todo: todoService, Router: router}
	t := todoHandler{s: todoService, logger: logger}
	e := errorHandler{}
	h := healthHandler{db: db}

	// redirect to list view
	s.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/todo?offset=0&limit=5", 301)
	}).Methods("GET")

	s.Router.HandleFunc("/todo/{id}", t.FindById).Methods("GET")
	s.Router.HandleFunc("/todo", t.FindAll).Methods("GET")
	s.Router.HandleFunc("/todo", t.Create).Methods("POST")
	s.Router.HandleFunc("/todo/{id}", t.Update).Methods("PUT")
	s.Router.HandleFunc("/todo/{id}", t.Done).Methods("PATCH")
	s.Router.HandleFunc("/todo/{id}", t.Delete).Methods("DELETE")

	s.Router.HandleFunc("/error", e.Error).Methods("POST")

	s.Router.HandleFunc("/health/live", h.Live).Methods("GET")
	s.Router.HandleFunc("/health/ready", h.Ready).Methods("GET")

	s.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

package controller

import (
	"log"
	"net/http"
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
	log.Printf("FindAll")
}

func (t *todoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Printf("Delete")
}

func (t *todoHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Printf("Update")
}

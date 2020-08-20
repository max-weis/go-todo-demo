package todo

import (
	"log"
	"todo"
)

type Service interface {
	Create(title, description string) (gotodo.Todo, error)
	FindById(id uint) (gotodo.Todo, error)
	FindAll(limit, offset int) ([]gotodo.Todo, error)
	Delete(id uint) (gotodo.Todo, error)
	Update(title, description string, status bool) (gotodo.Todo, error)
}

type service struct {
	repository gotodo.Repository
}

func NewService(repository gotodo.Repository) *service {
	return &service{repository: repository}
}

func (s *service) Create(title, description string) (gotodo.Todo, error) {
	panic("implement me")
}

func (s *service) FindById(id uint) (gotodo.Todo, error) {
	panic("implement me")
}

func (s *service) FindAll(limit, offset int) ([]gotodo.Todo, error) {
	log.Printf("find todos with limit: %d and offset: %d", limit, offset)

	return s.repository.FindAll(limit, offset)
}

func (s *service) Delete(id uint) (gotodo.Todo, error) {
	panic("implement me")
}

func (s *service) Update(title, description string, status bool) (gotodo.Todo, error) {
	panic("implement me")
}

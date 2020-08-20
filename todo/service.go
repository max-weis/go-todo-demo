package todo

import (
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
	if len(title) <= 0 {
		return gotodo.Todo{}, gotodo.TitleEmptyErr
	}
	if len(title) <= 50 {
		return gotodo.Todo{}, gotodo.TitleSizeErr
	}

	return s.repository.Create(title, description)
}

func (s *service) FindById(id uint) (gotodo.Todo, error) {
	panic("implement me")
}

func (s *service) FindAll(limit, offset int) ([]gotodo.Todo, error) {
	return s.repository.FindAll(limit, offset)
}

func (s *service) Delete(id uint) (gotodo.Todo, error) {
	panic("implement me")
}

func (s *service) Update(title, description string, status bool) (gotodo.Todo, error) {
	panic("implement me")
}

package todo

import (
	"log"
	"todo"
)

type loggingService struct {
	next Service
}

func NewLoggingService(next Service) *loggingService {
	return &loggingService{next: next}
}

func (l *loggingService) Create(title, description string) (gotodo.Todo, error) {
	log.Printf("create to todo with title: %s and description: %s", title, description)

	return l.next.Create(title, description)
}

func (l *loggingService) FindById(id uint) (gotodo.Todo, error) {
	panic("implement me")
}

func (l *loggingService) FindAll(limit, offset int) ([]gotodo.Todo, error) {
	log.Printf("find todos with limit: %d and offset: %d", limit, offset)

	return l.next.FindAll(limit, offset)
}

func (l *loggingService) Delete(id uint) (gotodo.Todo, error) {
	panic("implement me")
}

func (l *loggingService) Update(title, description string, status bool) (gotodo.Todo, error) {
	panic("implement me")
}

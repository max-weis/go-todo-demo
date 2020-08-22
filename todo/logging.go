package todo

import (
	"go.uber.org/zap"
	"todo"
)

type loggingService struct {
	next   Service
	logger zap.Logger
}

func NewLoggingService(next Service, logger zap.Logger) *loggingService {
	return &loggingService{next: next, logger: logger}
}

func (l *loggingService) Create(title, description string) (gotodo.Todo, error) {
	l.logger.Info("create to todo", zap.String("title", title), zap.String("description", description))

	return l.next.Create(title, description)
}

func (l *loggingService) FindById(id int) (gotodo.Todo, error) {
	l.logger.Info("find todo", zap.Int("id", id))

	return l.next.FindById(id)
}

func (l *loggingService) FindAll(limit, offset int) ([]gotodo.Todo, error) {
	l.logger.Info("find todos", zap.Int("limit", limit), zap.Int("offset", offset))

	return l.next.FindAll(limit, offset)
}

func (l *loggingService) Delete(id int) (gotodo.Todo, error) {
	l.logger.Info("delete todo", zap.Int("id", id))

	return l.next.Delete(id)
}

func (l *loggingService) Update(title, description string, status bool) (gotodo.Todo, error) {
	panic("implement me")
}

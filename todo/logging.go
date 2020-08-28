package todo

import (
	"go.uber.org/zap"
)

type loggingService struct {
	next   Service
	logger zap.Logger
}

func NewLoggingService(next Service, logger zap.Logger) *loggingService {
	return &loggingService{next: next, logger: logger}
}

func (l *loggingService) Create(title, description string) (Todo, error) {
	l.logger.Info("create to todo", zap.String("title", title), zap.String("description", description))

	return l.next.Create(title, description)
}

func (l *loggingService) FindById(id int) (Todo, error) {
	l.logger.Info("find todo", zap.Int("id", id))

	return l.next.FindById(id)
}

func (l *loggingService) FindAll(limit, offset int) ([]Todo, error) {
	l.logger.Info("find todos", zap.Int("limit", limit), zap.Int("offset", offset))

	return l.next.FindAll(limit, offset)
}

func (l *loggingService) Delete(id int) (Todo, error) {
	l.logger.Info("delete todo", zap.Int("id", id))

	return l.next.Delete(id)
}

func (l *loggingService) Update(id int, title, description string, status bool) (Todo, error) {
	l.logger.Info("update todo", zap.Int("id", id), zap.String("title", title), zap.String("description", description), zap.Bool("status", status))

	return l.next.Update(id, title, description, status)
}

func (l *loggingService) Done(id int, status bool) (Todo, error) {
	l.logger.Info("updates status", zap.Int("id", id), zap.Bool("status", status))

	return l.next.Done(id, status)
}

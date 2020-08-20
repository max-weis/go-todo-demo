package sqlite

import (
	"github.com/jinzhu/gorm"
	"todo"
)

type TodoRepository interface {
	Create(title, description string) (gotodo.Todo, error)
	FindById(id uint) (gotodo.Todo, error)
	FindAll(limit, offset uint) ([]gotodo.Todo, error)
	Delete(id uint) (gotodo.Todo, error)
	Update(title, description string, status bool) (gotodo.Todo, error)
}

type todoRepository struct {
	db gorm.DB
}

func NewTodoRepository(db gorm.DB) *todoRepository {
	return &todoRepository{db: db}
}

func (t *todoRepository) Create(title, description string) (gotodo.Todo, error) {
	panic("implement me")
}

func (t *todoRepository) FindById(id uint) (gotodo.Todo, error) {
	panic("implement me")
}

func (t *todoRepository) FindAll(limit, offset uint) ([]gotodo.Todo, error) {
	panic("implement me")
}

func (t *todoRepository) Delete(id uint) (gotodo.Todo, error) {
	panic("implement me")
}

func (t *todoRepository) Update(title, description string, status bool) (gotodo.Todo, error) {
	panic("implement me")
}

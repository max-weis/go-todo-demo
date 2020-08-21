package sqlite

import (
	"github.com/jinzhu/gorm"
	"log"
	"todo"
)

type TodoRepository interface {
	Create(title, description string) (gotodo.Todo, error)
	FindById(id int) (gotodo.Todo, error)
	FindAll(limit, offset int) ([]gotodo.Todo, error)
	Delete(id int) (gotodo.Todo, error)
	Update(title, description string, status bool) (gotodo.Todo, error)
}

type todoRepository struct {
	db gorm.DB
}

func NewTodoRepository(db gorm.DB) *todoRepository {
	return &todoRepository{db: db}
}

func (t *todoRepository) Create(title, description string) (gotodo.Todo, error) {
	todo := gotodo.NewTodo(title, description)
	newTodo := t.db.Create(todo)

	return *todo, newTodo.Error
}

func (t *todoRepository) FindById(id int) (gotodo.Todo, error) {
	var todo gotodo.Todo

	first := t.db.Find(&todo, id)

	return todo, first.Error
}

func (t *todoRepository) FindAll(limit, offset int) ([]gotodo.Todo, error) {
	var todos []gotodo.Todo
	find := t.db.Limit(limit).Offset(offset).Find(&todos)

	log.Printf("found %d todos", len(todos))

	return todos, find.Error
}

func (t *todoRepository) Delete(id int) (gotodo.Todo, error) {
	panic("implement me")
}

func (t *todoRepository) Update(title, description string, status bool) (gotodo.Todo, error) {
	panic("implement me")
}

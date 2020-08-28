package sqlite

import (
	"github.com/jinzhu/gorm"
	gotodo "github.com/max-weis/go-todo-demo"
)

type TodoRepository interface {
	Create(title, description string) (gotodo.Todo, error)
	FindById(id int) (gotodo.Todo, error)
	FindAll(limit, offset int) ([]gotodo.Todo, error)
	Delete(id int) (gotodo.Todo, error)
	Update(id int, title, description string, status bool) (gotodo.Todo, error)
	Done(id int, status bool) (gotodo.Todo, error)
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

	return todos, find.Error
}

func (t *todoRepository) Delete(id int) (gotodo.Todo, error) {
	var todo gotodo.Todo
	delete := t.db.Unscoped().Delete(&todo, id)

	return todo, delete.Error
}

func (t *todoRepository) Update(id int, title, description string, status bool) (gotodo.Todo, error) {
	var todo gotodo.Todo

	t.db.First(&todo, id)

	todo.Title = title
	todo.Description = description
	todo.Status = status

	newTodo := t.db.Save(&todo)

	return todo, newTodo.Error
}

func (t *todoRepository) Done(id int, status bool) (gotodo.Todo, error) {
	var todo gotodo.Todo

	t.db.First(&todo, id)

	todo.Status = status

	newTodo := t.db.Save(&todo)

	return todo, newTodo.Error
}

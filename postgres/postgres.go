package postgres

import (
	"github.com/max-weis/go-todo-demo/todo"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(title, description string) (todo.Todo, error)
	FindById(id int) (todo.Todo, error)
	FindAll(limit, offset int) ([]todo.Todo, error)
	Delete(id int) (todo.Todo, error)
	Update(id int, title, description string, status bool) (todo.Todo, error)
	Done(id int, status bool) (todo.Todo, error)
}

type todoRepository struct {
	db gorm.DB
}

func NewTodoRepository(db gorm.DB) *todoRepository {
	return &todoRepository{db: db}
}

func (t *todoRepository) Create(title, description string) (todo.Todo, error) {
	todo := todo.NewTodo(title, description)
	newTodo := t.db.Create(todo)

	return *todo, newTodo.Error
}

func (t *todoRepository) FindById(id int) (todo.Todo, error) {
	var todo todo.Todo

	first := t.db.Find(&todo, id)

	return todo, first.Error
}

func (t *todoRepository) FindAll(limit, offset int) ([]todo.Todo, error) {
	var todos []todo.Todo
	find := t.db.Limit(limit).Offset(offset).Find(&todos)

	return todos, find.Error
}

func (t *todoRepository) Delete(id int) (todo.Todo, error) {
	var todo todo.Todo
	delete := t.db.Unscoped().Delete(&todo, id)

	return todo, delete.Error
}

func (t *todoRepository) Update(id int, title, description string, status bool) (todo.Todo, error) {
	var todo todo.Todo

	t.db.First(&todo, id)

	todo.Title = title
	todo.Description = description
	todo.Status = status

	newTodo := t.db.Save(&todo)

	return todo, newTodo.Error
}

func (t *todoRepository) Done(id int, status bool) (todo.Todo, error) {
	var todo todo.Todo

	t.db.First(&todo, id)

	todo.Status = status

	newTodo := t.db.Save(&todo)

	return todo, newTodo.Error
}

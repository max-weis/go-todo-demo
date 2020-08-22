package gotodo

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model  // creates ID, createdAt, modifiedAt, deletedAt
	Title       string
	Description string
	Status      bool `gorm:"default:false"`
}

func NewTodo(title string, description string) *Todo {
	return &Todo{Title: title, Description: description}
}

type Service interface {
	Create(title, description string) (Todo, error)
	FindById(id int) (Todo, error)
	FindAll(limit, offset int) ([]Todo, error)
	Delete(id int) (Todo, error)
	Update(id int, title, description string, status bool) (Todo, error)
	Done(id int, status bool) (Todo, error)
}

type Repository interface {
	Create(title, description string) (Todo, error)
	FindById(id int) (Todo, error)
	FindAll(limit, offset int) ([]Todo, error)
	Delete(id int) (Todo, error)
	Update(id int, title, description string, status bool) (Todo, error)
	Done(id int, status bool) (Todo, error)
}

var TitleEmptyErr = errors.New("title should not be empty")
var TitleSizeErr = errors.New("title length should be between 1 and 50")

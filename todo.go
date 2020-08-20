package todo

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

type Service interface {
	Create(title, description string) (Todo, error)
	FindById(id uint) (Todo, error)
	FindAll(limit, offset uint) ([]Todo, error)
	Delete(id uint) (Todo, error)
	Update(title, description string, status bool) (Todo, error)
}

type Repository interface {
	Create(title, description string) (Todo, error)
	FindById(id uint) (Todo, error)
	FindAll(limit, offset uint) ([]Todo, error)
	Delete(id uint) (Todo, error)
	Update(title, description string, status bool) (Todo, error)
}

var TitleEmptyErr = errors.New("title should not be empty")
var TitleSizeErr = errors.New("title length should be between 1 and 50")

package todo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTodo(t *testing.T) {
	title := "test_title"
	description := "test_description"
	status := false

	todo := NewTodo(title, description)

	assert.Equal(t, title, todo.Title)
	assert.Equal(t, description, todo.Description)
	assert.Equal(t, status, todo.Status)
}

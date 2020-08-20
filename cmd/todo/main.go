package main

import (
	"fmt"
	"todo"
)

func main() {
	todo := todo.NewTodo("test title", "test description")

	fmt.Printf("id: %d, title: %s, description: %s, createdAt: %s, updatedAt: %s, deletedAt: %s", todo.ID, todo.Title, todo.Description, todo.CreatedAt, todo.UpdatedAt, todo.DeletedAt)
}

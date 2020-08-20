package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	gotodo "todo"
	"todo/controller"
	"todo/sqlite"
	"todo/todo"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	log.Println("start server")

	db := initDB()
	defer db.Close()

	var todoRepository sqlite.TodoRepository
	var todoService todo.Service

	todoRepository = sqlite.NewTodoRepository(*db)
	todoService = todo.NewService(todoRepository)

	router := mux.NewRouter()
	srv := controller.NewServer(todoService, router)

	http.ListenAndServe("0.0.0.0:8080", srv)

	log.Println("listening on port 8080")
}

func initDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		panic("could not read sqlite db")
	}

	db.AutoMigrate(&gotodo.Todo{})

	log.Println("db initialized")
	return db
}

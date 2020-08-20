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

	log.Println("listening on port 8080")

	http.ListenAndServe("0.0.0.0:8080", srv)

}

func initDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		panic("could not read sqlite db")
	}

	// create table
	db.AutoMigrate(&gotodo.Todo{})
	// populate table
	db.Create(&gotodo.Todo{Model: gorm.Model{}, Title: "Praesent ligula risus, imperdiet at.", Description: "Proin aliquet augue eget enim commodo, ut scelerisque tortor pretium. Maecenas dignissim sagittis condimentum. Praesent semper, ipsum nec porttitor iaculis.", Status: false})
	db.Create(&gotodo.Todo{Model: gorm.Model{}, Title: "Curabitur dictum condimentum ex, ut.", Description: "Duis quis sapien lorem. Sed ut luctus dolor. Etiam volutpat velit velit, vel rhoncus eros vestibulum facilisis. Sed aliquet neque.", Status: true})
	db.Create(&gotodo.Todo{Model: gorm.Model{}, Title: "Nullam placerat ante vitae orci.", Description: "Nunc pretium enim malesuada quam gravida volutpat sed a sapien. Cras scelerisque diam sapien, quis consectetur felis euismod sed. Cras.", Status: false})
	db.Create(&gotodo.Todo{Model: gorm.Model{}, Title: "Fusce iaculis vel magna vel.", Description: "Phasellus scelerisque libero non nisi porttitor aliquet. Curabitur eu sapien at purus mollis condimentum. Aliquam ut odio augue. Aenean pretium.", Status: false})
	db.Create(&gotodo.Todo{Model: gorm.Model{}, Title: "Phasellus lacinia sollicitudin erat, a.", Description: "Etiam egestas malesuada augue et lacinia. Cras at eleifend ligula, nec accumsan justo. Aliquam eget dictum elit, id aliquam elit.", Status: true})
	db.Create(&gotodo.Todo{Model: gorm.Model{}, Title: "Vivamus accumsan diam et tortor.", Description: "Maecenas non tempus magna. Nam vehicula diam dui. Vestibulum non maximus nibh. Aenean est lacus, faucibus quis interdum quis, maximus.", Status: false})

	log.Println("db initialized")
	return db
}

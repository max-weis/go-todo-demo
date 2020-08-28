package main

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"

	gotodo "github.com/max-weis/go-todo-demo"
	"github.com/max-weis/go-todo-demo/controller"
	"github.com/max-weis/go-todo-demo/postgres"
	"github.com/max-weis/go-todo-demo/todo"
	"net/http"

	pq "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var logger zap.Logger

func main() {
	initLogger()
	defer logger.Sync()

	logger.Info("start server")

	db := initDB()

	var todoRepository postgres.TodoRepository
	var todoService todo.Service

	todoRepository = postgres.NewTodoRepository(*db)
	todoService = todo.NewService(todoRepository)
	todoService = todo.NewLoggingService(todoService, logger)

	router := mux.NewRouter()
	srv := controller.NewServer(todoService, logger, router)

	logger.Info("listening on port 8080")

	http.ListenAndServe("0.0.0.0:8080", srv)

}

func initLogger() {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./log/monitoring.log",
		"stdout",
	}
	build, err := cfg.Build()
	if err != nil {
		panic("cannot init logger")
	}

	logger = *build
}

func initDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(pq.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("error", zap.Error(err))
		panic("could not read postgres db")
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

	logger.Info("db initialized")
	return db
}

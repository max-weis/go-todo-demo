package main

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap"

	"github.com/max-weis/go-todo-demo/postgres"
	"github.com/max-weis/go-todo-demo/server"
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

	// definition at the bottom
	todoService = todo.NewInstrumentingService(createCounter, createLatency, findOneCounter, findOneLatency, findAllCounter, findAllLatency, deleteCounter, deleteLatency, updateCounter, updateLatency, doneCounter, doneLatency, todoService)

	router := mux.NewRouter()
	srv := server.NewServer(todoService, logger, router, db)

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
	config := NewConfig()

	db, err := gorm.Open(pq.Open(config.getDSN()), &gorm.Config{})
	if err != nil {
		logger.Error("error", zap.Error(err))
		panic("could not read postgres db")
	}

	// create table
	db.AutoMigrate(&todo.Todo{})
	// populate table
	db.Create(&todo.Todo{Model: gorm.Model{}, Title: "Praesent ligula risus, imperdiet at.", Description: "Proin aliquet augue eget enim commodo, ut scelerisque tortor pretium. Maecenas dignissim sagittis condimentum. Praesent semper, ipsum nec porttitor iaculis.", Status: false})
	db.Create(&todo.Todo{Model: gorm.Model{}, Title: "Curabitur dictum condimentum ex, ut.", Description: "Duis quis sapien lorem. Sed ut luctus dolor. Etiam volutpat velit velit, vel rhoncus eros vestibulum facilisis. Sed aliquet neque.", Status: true})
	db.Create(&todo.Todo{Model: gorm.Model{}, Title: "Nullam placerat ante vitae orci.", Description: "Nunc pretium enim malesuada quam gravida volutpat sed a sapien. Cras scelerisque diam sapien, quis consectetur felis euismod sed. Cras.", Status: false})
	db.Create(&todo.Todo{Model: gorm.Model{}, Title: "Fusce iaculis vel magna vel.", Description: "Phasellus scelerisque libero non nisi porttitor aliquet. Curabitur eu sapien at purus mollis condimentum. Aliquam ut odio augue. Aenean pretium.", Status: false})
	db.Create(&todo.Todo{Model: gorm.Model{}, Title: "Phasellus lacinia sollicitudin erat, a.", Description: "Etiam egestas malesuada augue et lacinia. Cras at eleifend ligula, nec accumsan justo. Aliquam eget dictum elit, id aliquam elit.", Status: true})
	db.Create(&todo.Todo{Model: gorm.Model{}, Title: "Vivamus accumsan diam et tortor.", Description: "Maecenas non tempus magna. Nam vehicula diam dui. Vestibulum non maximus nibh. Aenean est lacus, faucibus quis interdum quis, maximus.", Status: false})

	logger.Info("db initialized")
	return db
}

var (
	createCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "create_total",
			Help: "The total number of created todos",
		},
	)
	createLatency = promauto.NewSummary(
		prometheus.SummaryOpts{
			Name: "create_latency",
			Help: "Total duration of create todo requests",
		},
	)
	findOneCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "find_one_total",
			Help: "The total number of found todos",
		},
	)
	findOneLatency = promauto.NewSummary(
		prometheus.SummaryOpts{
			Name: "find_one_latency",
			Help: "Total duration of find one todo requests",
		},
	)
	findAllCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "find_all_total",
			Help: "The total number of all found todos",
		},
	)
	findAllLatency = promauto.NewSummary(
		prometheus.SummaryOpts{
			Name: "find_all_latency",
			Help: "Total duration of find all todo requests",
		},
	)
	deleteCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "delete_total",
			Help: "The total number of deleted todos",
		},
	)
	deleteLatency = promauto.NewSummary(
		prometheus.SummaryOpts{
			Name: "delete_latency",
			Help: "Total duration of deleted todo requests",
		},
	)
	updateCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "update_total",
			Help: "The total number of updated todos",
		},
	)
	updateLatency = promauto.NewSummary(
		prometheus.SummaryOpts{
			Name: "update_latency",
			Help: "Total duration of updated todo requests",
		},
	)
	doneCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "done_total",
			Help: "The total number of done todos",
		},
	)
	doneLatency = promauto.NewSummary(
		prometheus.SummaryOpts{
			Name: "done_latency",
			Help: "Total duration of done todo requests",
		},
	)
)

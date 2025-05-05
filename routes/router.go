package routes

import (
	"todo-app/controller" // go.mod'daki module adı buysa doğru olur

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/todos", controller.CreateTodo).Methods("POST")
	router.HandleFunc("/todos", controller.GetTodos).Methods("GET")

	return router
}

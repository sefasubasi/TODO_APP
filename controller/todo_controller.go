// controller/todo_controller.go
package Controller

import (
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login endpoint"))
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create todo endpoint"))
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get todos endpoint"))
}

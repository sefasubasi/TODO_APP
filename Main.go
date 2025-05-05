package main

import (
	"log"
	"net/http"
	"todo-app/routes" // modül adıyla aynı olmalı
)

func main() {
	router := routes.NewRouter()
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", router)
}

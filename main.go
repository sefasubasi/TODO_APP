package main

import (
	"todo_app/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}

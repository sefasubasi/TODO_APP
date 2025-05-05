// model/todo.go
package model

import "time"

type TodoStep struct {
	ID        int
	Content   string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type TodoList struct {
	ID        int
	Name      string
	Steps     []TodoStep
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

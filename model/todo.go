package model

import "time"

type TodoList struct {
	ID        int
	UserID    int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Steps     []TodoStep
}

type TodoStep struct {
	ID        int
	ListID    int
	Content   string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

package mock

import "todo_app/model"

var Users = []model.User{
	{ID: 1, Username: "user1", Password: "123456", Role: model.RoleUser},
	{ID: 2, Username: "admin1", Password: "admin123", Role: model.RoleAdmin},
}

var TodoLists = []model.TodoList{}
var TodoSteps = []model.TodoStep{}

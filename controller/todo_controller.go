package controller

import (
	"net/http"
	"strconv"
	"time"
	"todo_app/mock"
	"todo_app/model"

	"github.com/gin-gonic/gin"
)

func GetOwnTodos(c *gin.Context) {
	userID := c.GetInt("userID")

	var todos []model.TodoList
	for _, todo := range mock.TodoLists {
		if todo.UserID == userID && todo.DeletedAt == nil {
			todos = append(todos, todo)
		}
	}

	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	userID := c.GetInt("userID")

	var input struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri"})
		return
	}

	now := time.Now()
	newTodo := model.TodoList{
		ID:        len(mock.TodoLists) + 1,
		UserID:    userID,
		Name:      input.Name,
		CreatedAt: now,
		UpdatedAt: now,
		Steps:     []model.TodoStep{},
	}

	mock.TodoLists = append(mock.TodoLists, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

func UpdateTodo(c *gin.Context) {
	userID := c.GetInt("userID")
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID"})
		return
	}

	var input struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri"})
		return
	}

	for i, todo := range mock.TodoLists {
		if todo.ID == id && todo.DeletedAt == nil {
			if todo.UserID != userID {
				c.JSON(http.StatusForbidden, gin.H{"error": "Bu liste size ait değil"})
				return
			}

			mock.TodoLists[i].Name = input.Name
			mock.TodoLists[i].UpdatedAt = time.Now()

			c.JSON(http.StatusOK, mock.TodoLists[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Liste bulunamadı"})
}

func DeleteTodo(c *gin.Context) {
	userID := c.GetInt("userID")
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID"})
		return
	}

	for i, todo := range mock.TodoLists {
		if todo.ID == id && todo.DeletedAt == nil {
			if todo.UserID != userID {
				c.JSON(http.StatusForbidden, gin.H{"error": "Bu liste size ait değil"})
				return
			}

			now := time.Now()
			mock.TodoLists[i].DeletedAt = &now

			c.JSON(http.StatusOK, gin.H{"message": "Liste silindi (soft delete)"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Liste bulunamadı"})

}

func GetAllTodos(c *gin.Context) {
	var todos []model.TodoList
	for _, todo := range mock.TodoLists {
		if todo.DeletedAt == nil {
			todos = append(todos, todo)
		}
	}

	c.JSON(http.StatusOK, todos)
}

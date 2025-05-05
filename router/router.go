package router

import (
	"github.com/gin-gonic/gin"

	"todo_app/controller"
	"todo_app/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Giriş rotası
	r.POST("/login", controller.Login)

	// TO-DO işlemleri (Giriş yapılmış kullanıcılar için)
	todo := r.Group("/todo")
	todo.Use(middleware.JWTAuthMiddleware())
	{
		todo.GET("/my", controller.GetOwnTodos)
		todo.POST("/", controller.CreateTodo)
		todo.PUT("/:id", controller.UpdateTodo)
		todo.DELETE("/:id", controller.DeleteTodo)
	}

	// Admin işlemleri (Tüm kullanıcıların TO-DO verileri)
	admin := r.Group("/admin")
	admin.Use(middleware.JWTAuthMiddleware(), middleware.AdminOnly())
	{
		admin.GET("/todos", controller.GetAllTodos)
	}

	return r
}

package handlers

import (
	"github.com/phucnh/api-mock/src/app"
)

func AddRouters(a *app.App) {
	// Authentication handlers
	userHandler := NewUserHandler()
	a.POST("/signup", userHandler.Signup)
	a.GET("/login", userHandler.ValidateUser)
	a.POST("/login", userHandler.Login)
	a.GET("/profile", userHandler.GetProfile)
	a.POST("/change_password", userHandler.ChangePassword)
	a.POST("/logout", userHandler.Logout)

	// User handlers
	a.GET("/users", userHandler.List)
	a.GET("/users/{id}", userHandler.Info)
	a.PUT("/users/{id}", userHandler.Update)
	a.DELETE("/users/{id}", userHandler.Delete)

	// Todo handlers
	todoHandler := NewTodoHander()
	a.GET("/todos", todoHandler.List)
	a.POST("/todos", todoHandler.Create)
}

package handlers

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pa-vuhn/api-mock/src/app"
	"github.com/pa-vuhn/api-mock/src/config"
	"github.com/pa-vuhn/api-mock/src/entities"
	"github.com/pa-vuhn/api-mock/src/services"
)

type (
	TodoHandler struct {
		todoServices *services.TodoService
	}
)

func NewTodoHander() *TodoHandler {
	todoServices := services.NewTodoSerivce()
	return &TodoHandler{
		todoServices: todoServices,
	}
}

func (h *TodoHandler) List(c app.Context) {
	userID := c.FormValue("user_id")

	todos, err := h.todoServices.List(c, userID)
	if err != nil {
		renderError(c, err)
		return
	}

	resp := map[string]interface{}{
		"data": todos,
	}

	c.RenderJSON(resp)
}

func (h *TodoHandler) Create(c app.Context) {
	todo := &entities.Todo{}
	err := c.BindJSON(todo)
	if err != nil {
		renderError(c, err)
		return
	}
	userID := c.ContextValue(config.UserAuthCtxKey)
	todo.UserID = fmt.Sprintf("%v", userID)
	todo.ID = uuid.New().String()

	todo, err = h.todoServices.Create(c, todo)
	if err != nil {
		renderError(c, err)
		return
	}

	data := map[string]*entities.Todo{
		"data": todo,
	}
	c.RenderJSON(data)
}

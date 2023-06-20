package services

import (
	"log"

	"github.com/phucnh/api-mock/src/app"
	"github.com/phucnh/api-mock/src/config"
	"github.com/phucnh/api-mock/src/entities"
	"github.com/phucnh/api-mock/src/models"
)

type (
	TodoService struct {
		todoModel models.TodoModel
		userModel models.UserModel
	}
)

func NewTodoSerivce() *TodoService {
	todoModel := models.NewTodoModel()
	userModel := models.NewUserModel()
	return &TodoService{
		todoModel: todoModel,
		userModel: userModel,
	}
}

func (s *TodoService) List(c app.Context, userID string) ([]*entities.Todo, error) {
	todos, err := s.todoModel.List(c, userID)
	if err != nil {
		return nil, config.ErrorServerError
	}

	return todos, nil
}

func (s *TodoService) Create(c app.Context, todo *entities.Todo) (*entities.Todo, error) {
	err := s.todoModel.Create(c, todo)
	if err != nil {
		log.Println(err)
		return nil, config.ErrorServerError
	}

	return todo, nil
}

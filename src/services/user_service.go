package services

import (
	"log"

	"github.com/phucnh/api-mock/src/config"
	"github.com/phucnh/api-mock/src/entities"
	"github.com/phucnh/api-mock/src/utils/jwt"

	"github.com/phucnh/api-mock/src/app"
	"github.com/phucnh/api-mock/src/models"
)

type (
	UserService struct {
		userModel models.UserModel
	}
)

func NewUserService() *UserService {
	userModel := models.NewUserModel()
	return &UserService{
		userModel: userModel,
	}
}

func (s *UserService) List(c app.Context) ([]*entities.User, error) {
	users, err := s.userModel.GetUsersList(c)
	if err != nil {
		log.Println(err)
		return nil, config.ErrorServerError
	}

	return users, nil
}

func (s *UserService) Info(c app.Context, userID string) (*entities.User, error) {
	user, err := s.userModel.GetUserByID(c, userID)
	if err != nil {
		log.Println(err)
		return nil, config.ErrorServerError
	}

	return user, nil
}

func (s *UserService) Create(c app.Context, user *entities.User) error {
	if err := s.userModel.CreateUser(c, user); err != nil {
		log.Println(err)
		return config.ErrorServerError
	}

	return nil
}

func (s *UserService) Update(c app.Context, user *entities.User) error {
	if err := s.userModel.UpdateUser(c, user); err != nil {
		log.Println(err)
		return config.ErrorServerError
	}

	return nil
}

func (s *UserService) Delete(c app.Context, userID string) error {
	if err := s.userModel.DeleteUser(c, userID); err != nil {
		log.Println(err)
		return config.ErrorServerError
	}

	return nil
}

func (s *UserService) ChangePassword(c app.Context, newPassword, userID string) error {
	if err := s.userModel.ChangePassword(c, newPassword, userID); err != nil {
		log.Println(err)
		return config.ErrorServerError
	}

	return nil
}

func (t *UserService) ValidateUser(c app.Context, email, password string) (string, error) {
	user, err := t.userModel.ValidateUser(c, email, password)
	if err != nil {
		return "", config.ErrorLoginError
	}
	token, err := jwt.CreateUserToken(user.ID, c.GetEnv("JWT_KEY"))
	if err != nil {
		return "", config.ErrorServerError
	}
	return token, nil
}

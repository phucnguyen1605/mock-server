package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/google/uuid"
	"github.com/phucnh/api-mock/src/app"
	"github.com/phucnh/api-mock/src/config"
	"github.com/phucnh/api-mock/src/entities"
	"github.com/phucnh/api-mock/src/services"
)

type (
	UserHandler struct {
		userService *services.UserService
	}

	changePasswordParam struct {
		NewPassword string `json:"new_password"`
	}

	loginParam struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	userProfile struct {
		ID       string `json:"id"`
		Email    string `json:"email"`
		Fullname string `json:"fullname"`
	}
)

func NewUserHandler() *UserHandler {
	userService := services.NewUserService()
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) List(c app.Context) {
	users, err := h.userService.List(c)
	if err != nil {
		renderError(c, err)
		return
	}

	resp := map[string]interface{}{
		"data": users,
	}

	c.RenderJSON(resp)
}

func (h *UserHandler) Info(c app.Context) {
	userID := mux.Vars(c.Request())["id"]
	user, err := h.userService.Info(c, userID)
	if err != nil {
		renderError(c, err)
		return
	}
	data := map[string]*entities.User{
		"data": user,
	}
	c.RenderJSON(data)
}

func (h *UserHandler) Update(c app.Context) {
	user := &entities.User{}
	err := c.BindJSON(user)
	if err != nil {
		renderError(c, err)
		return
	}
	userID := mux.Vars(c.Request())["id"]
	user.ID = userID
	err = h.userService.Update(c, user)
	if err != nil {
		renderError(c, err)
		return
	}
	data := map[string]*entities.User{
		"data": user,
	}
	c.RenderJSON(data)
}

func (h *UserHandler) Delete(c app.Context) {
	userID := mux.Vars(c.Request())["id"]

	if err := h.userService.Delete(c, userID); err != nil {
		renderError(c, err)
		return
	}
	c.RenderEmptyBody(http.StatusNoContent)
}

func (h *UserHandler) Signup(c app.Context) {
	user := &entities.User{}
	err := c.BindJSON(user)
	if err != nil {
		renderError(c, err)
		return
	}
	user.ID = uuid.New().String()

	err = h.userService.Create(c, user)
	if err != nil {
		renderError(c, err)
		return
	}

	data := map[string]*entities.User{
		"data": user,
	}
	c.RenderJSON(data)
}

func (h *UserHandler) ValidateUser(c app.Context) {
	id := c.FormValue("email")
	password := c.FormValue("password")

	token, err := h.userService.ValidateUser(c, id, password)
	if err != nil {
		renderError(c, err)
		return
	}
	data := map[string]string{
		"data": token,
	}
	c.RenderJSON(data)
}

func (h *UserHandler) Login(c app.Context) {
	loginInfo := &loginParam{}
	err := c.BindJSON(loginInfo)
	if err != nil {
		renderError(c, err)
		return
	}
	log.Println(loginInfo.Email, loginInfo.Password)

	token, err := h.userService.ValidateUser(c, loginInfo.Email, loginInfo.Password)
	if err != nil {
		renderError(c, err)
		return
	}
	data := map[string]string{
		"data": token,
	}
	c.RenderJSON(data)
}

func (h *UserHandler) GetProfile(c app.Context) {
	userID := c.ContextValue(config.UserAuthCtxKey)
	user, err := h.userService.Info(c, fmt.Sprintf("%v", userID))
	if err != nil {
		renderError(c, err)
		return
	}
	data := map[string]*userProfile{
		"data": &userProfile{
			ID:       user.ID,
			Fullname: user.Fullname,
			Email:    user.Email,
		},
	}
	c.RenderJSON(data)
}

func (h *UserHandler) ChangePassword(c app.Context) {
	p := &changePasswordParam{}
	if err := c.BindJSON(p); err != nil {
		renderError(c, err)
		return
	}
	log.Println(p.NewPassword)

	userID := c.ContextValue(config.UserAuthCtxKey)
	if err := h.userService.ChangePassword(c, p.NewPassword, fmt.Sprintf("%v", userID)); err != nil {
		renderError(c, err)
		return
	}

	data := map[string]bool{
		"data": true,
	}
	c.RenderJSON(data)
}

func (h *UserHandler) Logout(c app.Context) {
	data := map[string]bool{
		"data": true,
	}
	c.RenderJSON(data)
}

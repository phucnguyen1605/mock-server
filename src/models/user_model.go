package models

import (
	"log"

	"github.com/phucnh/api-mock/src/app"
	"github.com/phucnh/api-mock/src/entities"
)

type (
	UserModel interface {
		CreateUser(c app.Context, p *entities.User) error
		UpdateUser(c app.Context, user *entities.User) error
		ValidateUser(c app.Context, userID, password string) (*entities.User, error)
		GetUserByID(c app.Context, userID string) (*entities.User, error)
		ChangePassword(c app.Context, userID, newPassword string) error
		GetUsersList(c app.Context) ([]*entities.User, error)
		DeleteUser(c app.Context, userID string) error
	}

	sqlUserModel struct{}
)

func NewUserModel() UserModel {
	return &sqlUserModel{}
}

func (m *sqlUserModel) CreateUser(c app.Context, p *entities.User) error {
	sql := `INSERT INTO users (id, email, fullname, password) VALUES (?, ?, ?, ?)`
	return c.DB().Exec(sql, p.ID, p.Email, p.Fullname, p.Password)
}

func (m *sqlUserModel) ValidateUser(c app.Context, email, password string) (*entities.User, error) {
	sql := `SELECT id FROM users WHERE email = ? AND password = ?`
	user := new(entities.User)
	if err := c.DB().QueryRow(user, sql, email, password); err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

func (m *sqlUserModel) GetUsersList(c app.Context) ([]*entities.User, error) {
	sqlStr := `SELECT * FROM users`
	rows, err := c.DB().Query(sqlStr)
	if err != nil {
		return nil, err
	}

	users := []*entities.User{}
	for rows.Next() {
		var user = &entities.User{}
		rows.StructScan(user)
		users = append(users, user)
	}

	return users, nil
}

func (m *sqlUserModel) GetUserByID(c app.Context, userID string) (*entities.User, error) {
	sql := `SELECT * FROM users WHERE id = ?`
	user := new(entities.User)
	if err := c.DB().QueryRow(user, sql, userID); err != nil {
		return nil, err
	}

	return user, nil
}

func (m *sqlUserModel) ChangePassword(c app.Context, newPassword, userID string) error {
	sql := `UPDATE users SET password = ? WHERE id = ?`
	return c.DB().Exec(sql, newPassword, userID)
}

func (m *sqlUserModel) UpdateUser(c app.Context, user *entities.User) error {
	sql := `UPDATE users SET email =?, fullname = ? WHERE id = ?`
	return c.DB().Exec(sql, user.Email, user.Fullname, user.ID)
}

func (m *sqlUserModel) DeleteUser(c app.Context, userID string) error {
	sql := `DELETE FROM users WHERE id = ?`
	return c.DB().Exec(sql, userID)
}

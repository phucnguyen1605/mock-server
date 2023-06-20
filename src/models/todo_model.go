package models

import (
	"log"

	"github.com/phucnh/api-mock/src/app"
	"github.com/phucnh/api-mock/src/entities"
)

type (
	TodoModel interface {
		Create(c app.Context, t *entities.Todo) error
		List(c app.Context, userID string) ([]*entities.Todo, error)
		CountTodoByUserPerDay(c app.Context, userID, date string) (int, error)
	}

	pgTodoModel struct{}

	countRow struct {
		Total int `db:"total"`
	}
)

func NewTodoModel() TodoModel {
	return &pgTodoModel{}
}

func (m *pgTodoModel) Create(c app.Context, t *entities.Todo) error {
	sql := `INSERT INTO todos (id, user_id, title, description) 
		VALUES (?, ?, ?, ?)`
	return c.DB().Exec(sql, t.ID, t.UserID, t.Title, t.Description)
}

func (m *pgTodoModel) List(c app.Context, userID string) ([]*entities.Todo, error) {
	sqlStr := `SELECT 
			id,
			user_id,
			title,
			description
		FROM todos`
	rows, err := c.DB().Query(sqlStr)
	if err != nil {
		return nil, err
	}

	todos := []*entities.Todo{}
	for rows.Next() {
		var todo = &entities.Todo{}
		rows.StructScan(todo)
		todos = append(todos, todo)
	}

	return todos, nil
}

func (m *pgTodoModel) CountTodoByUserPerDay(c app.Context, userID, date string) (int, error) {
	row := &countRow{}
	sql := "SELECT COUNT(*) AS total FROM todos WHERE user_id = ? AND created_at = ?"
	if err := c.DB().QueryRow(row, sql, userID, date); err != nil {
		log.Println(err)
		return 0, err
	}

	return row.Total, nil
}

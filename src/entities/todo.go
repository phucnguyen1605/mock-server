package entities

type (
	Todo struct {
		ID          string `db:"id" json:"id"`
		UserID      string `db:"user_id" json:"user_id"`
		Title       string `db:"title" json:"title"`
		Description string `db:"description" json:"description"`
		CreatedAt   string `db:"created_at" json:"created_at"`
	}
)

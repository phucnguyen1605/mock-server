package entities

// User reflects users data from DB
type User struct {
	ID       string `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Fullname string `db:"fullname" json:"fullname"`
	Password string `db:"password" json:"password"`
}

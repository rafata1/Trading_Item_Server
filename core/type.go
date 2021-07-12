package core

type User struct {
	ID int `db:"id"`
	Email string `db:"email"`
	Username string `db:"username"`
	Password string `db:"password"`
}

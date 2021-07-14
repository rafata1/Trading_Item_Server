package core

type User struct {
	ID int `db:"id"`
	Email string `db:"email"`
	Username string `db:"username"`
	Password string `db:"password"`
	PhoneNumber string `db:"phone_number"`
	Gender string `db:"gender"`
	DateOfBirth string `db:"dob"`
}

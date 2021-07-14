package core

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func connectToDB() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "sql6425032:bnMVGBV6SF@(sql6.freesqldatabase.com:3306)/sql6425032")
	if err != nil {
		log.Fatalf("Err occured while connecting to db")
		return &sqlx.DB{}
	}
	fmt.Println("Connect to db successfully")
	return db
}

func signUp(email string, username string, password string, phone_number string,
	gender string, dob string ,myDB *sqlx.DB) (statusCode int32,
	detail string) {

	//check if email is already existed
	var user User
	myDB.Get(&user, "SELECT id, email, username, password FROM Users WHERE email=? LIMIT 1", email)

	if user.Email == email {
		return 101, "Email đã tồn tại"
	}

	myDB.Get(&user, "SELECT id, email, username, password FROM Users WHERE username=? LIMIT 1", username)

	if user.Username == username {
		return 202, "Username đã tồn tại"
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	queryCode := `INSERT INTO Users (email, username, password, phone_number, gender, dob) VALUES (?, ?, ?, ?, ?, ?)`
	myDB.MustExec(queryCode, email, username, string(hashedPassword), phone_number, gender, dob)
	return 200, "Thêm tài khoản thành công"
}

func login(email string, password string, myDB *sqlx.DB) (statusCode int32, detail string) {

	//check if email is existed
	var user User
	myDB.Get(&user, "SELECT id, email, username, password, phone_number, gender, dob FROM Users WHERE email=? LIMIT 1", email)
	if user.Email != email {
		return 101, "Email không tồn tại"
	}
	//check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return 202, "Password sai"
	}

	return 200, "Đăng nhập thành công"
}


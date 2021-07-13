package core

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func connectToDB() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "root:1@(mysql-server:3306)/trading_item_db")
	if err != nil {
		log.Fatalf("Err occured while connecting to db")
		return &sqlx.DB{}
	}
	fmt.Println("Connect to db successfully")
	return db
}

func addUser(email string, username string, password string, myDB *sqlx.DB) (statusCode int32, detail string) {

	//check if email is already existed
	var user User
	myDB.Get(&user, "SELECT id, email, username, password FROM Users WHERE email=? LIMIT 1", email)

	if user.Email == email {
		return 101, "Email da ton tai"
	}

	myDB.Get(&user, "SELECT id, email, username, password FROM Users WHERE username=? LIMIT 1", username)

	if user.Username == username {
		return 202, "Username da ton tai"
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	queryCode := `INSERT INTO Users (email, username, password) VALUES (?, ?, ?)`
	myDB.MustExec(queryCode, email, username, string(hashedPassword))
	return 200, "Thêm tài khoản thành công"
}

func login(email string, password string, myDB *sqlx.DB) (statusCode int32, detail string) {

	//check if email is existed
	var user User
	myDB.Get(&user, "SELECT id, email, username, password FROM Users WHERE email=? LIMIT 1", email)

	if user.Email != email {
		return 101, "Email khong ton tai"
	}

	//check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return 202, "Password sai"
	}

	return 200, "Dang nhap thanh cong"
}


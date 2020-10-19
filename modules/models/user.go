package usermodel

import (
	"database/sql"
	"fmt"
	"log"

	// for mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	selectAllUserQuery = "SELECT id, user_email, user_password, user_name FROM user_master"
	loginUserQuery     = "SELECT id, user_email, user_password, user_name FROM user_master WHERE user_name='%s' AND  user_password='%s'"
)

// User - is a interface for user table in database
type User struct {
	ID       int    `json:"id"`
	Username string `json:"user_name"`
	Usermail string `json:"user_email"`
	Password string `json:"user_password"`
}

// GetAllUsers - is a function to return All users from database
func GetAllUsers(db *sql.DB) (users []User) {
	results, err := db.Query(selectAllUserQuery)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {

		var user User
		err = results.Scan(&user.ID, &user.Username, &user.Usermail, &user.Password)
		if err != nil {
			panic(err.Error())
		}

		log.Println(user.ID)
		log.Printf(user.Username)
		log.Printf(user.Usermail)
		log.Printf(user.Password)
		users = append(users, user)
	}
	return
}

// LoginResult - is a function to return LOGIN result from database
func LoginResult(db *sql.DB, user string, pass string) (users []User) {
	query := fmt.Sprintf(loginUserQuery, user, pass)
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {

		var user User
		err = results.Scan(&user.ID, &user.Username, &user.Usermail, &user.Password)
		if err != nil {
			panic(err.Error())
		}

		log.Println(user.ID)
		log.Printf(user.Username)
		log.Printf(user.Usermail)
		log.Printf(user.Password)
		users = append(users, user)
	}
	return
}

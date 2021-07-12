package userdata

import (
	"database/sql"
	"fmt"
	types "resful-server/datastruct"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host      = "mysql"
	database  = "testdb"
	user      = "root"
	password  = "root"
	tableName = "Users"
)

func dbOpen() *sql.DB {
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)
	conn, err := sql.Open("mysql", connectionString)
	checkError(err)
	// Lượng truy cập tối đa cho phép đến DataBase, mặc định là 150
	conn.SetMaxOpenConns(3000)
	conn.SetMaxIdleConns(2000)
	return conn
}

func SelectPasswords(username string) types.User {
	db := dbOpen()
	defer db.Close()
	password, err := db.Query("SELECT * FROM Users WHERE username=?", username)
	checkError(err)
	defer password.Close()

	var userData types.User
	for password.Next() {
		err = password.Scan(&userData.ID, &userData.Username, &userData.Passwords, &userData.Firstname, &userData.Lastname)
		if err != nil {
			panic(err)
			return userData
		}
	}
	return userData
}

func InsertUser(user types.User) string {
	db := dbOpen()
	defer db.Close()
	insert, err := db.Prepare("insert into Users (username, passwords, firstname, lastname) value (?, ?, ?, ?)")
	checkError(err)
	defer insert.Close()

	_, errExec := insert.Exec(user.Username, user.Passwords, user.Firstname, user.Lastname)
	if errExec != nil {
		return "Register failure, please check your username!"
	} else {
		return "Register successfully"
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

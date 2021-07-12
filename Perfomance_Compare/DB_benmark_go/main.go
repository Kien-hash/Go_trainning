package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
}

var db *sql.DB
var err error
var tableName string = "users"

func main() {
	// Open connect to DataBase
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/testdb")

	if err != nil {
		panic(err.Error())
	}

	// Lượng truy cập tối đa cho phép đến DataBase, mặc định là 150
	db.SetMaxOpenConns(3000)
	db.SetMaxIdleConns(2000)

	defer db.Close()

	// The API callback list here
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/read", readTable).Methods("GET")
	router.HandleFunc("/insert", insertTable).Methods("POST")
	router.HandleFunc("/update", updateTable).Methods("PUT")
	router.HandleFunc("/delete/{id}", deleteTable).Methods("DELETE")
	fmt.Println("Server start at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func insertTable(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/json")
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		panic(err.Error())
	} else {
		// fmt.Println(user.ID, user.Name)
	}

	s := "INSERT INTO " + tableName + " VALUES ( " + strconv.Itoa(user.ID) + ", '" + user.Name + "' )"
	// fmt.Println(s)

	// Query insert to database
	insert, err := db.Query(s)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	fmt.Fprintf(response, "Table ["+tableName+"] was inserted!")
}

func readTable(response http.ResponseWriter, request *http.Request) {
	var Users []User

	// Query insert to database
	read, err := db.Query("SELECT * FROM users")

	// if there is an error inserting, handle it
	if err != nil {
		//panic(err.Error())
	}

	defer read.Close()

	// console log the result from read
	for read.Next() {
		var user User

		err = read.Scan(&user.ID, &user.Name)
		if err != nil {
			panic(err.Error())
		}
		Users = append(Users, user)
		// fmt.Println(user.ID, ": ", user.Name)
	}
	json.NewEncoder(response).Encode(Users)

}

func createTable(table_name string) {
	// Query insert to database
	create, err := db.Query("Create table " + table_name + " (ID int(11), Name Char(64));")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer create.Close()
}

func updateTable(response http.ResponseWriter, request *http.Request) {
	// change values in a single row .
	// Query insert to database

	response.Header().Set("content-type", "application/json")
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		panic(err.Error())
	} else {
		// fmt.Println(user.ID, user.Name)
	}

	//fmt.Println(temp.ID,": ",temp.Name)
	update, err := db.Query("UPDATE " + tableName + " SET name = '" + user.Name + "' WHERE ID = " + strconv.Itoa(user.ID))
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer update.Close()
	fmt.Fprintln(response, "ID = "+strconv.Itoa(user.ID)+" is updated!")

}

func deleteTable(response http.ResponseWriter, request *http.Request) {
	// Query insert to database
	vars := mux.Vars(request)
	key := vars["id"]

	if err != nil {
		panic(err.Error())
	} else {
		// fmt.Println(user.ID, user.Name)
	}

	del, err := db.Query("DELETE FROM `users` WHERE `ID` =  " + key)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer del.Close()
	//fmt.Fprintln(response, "ID = "+strconv.Itoa(user.ID)+" is deleted!")
}

package database

import (
	"database/sql"
	"fmt"
	"sort"
	"strconv"
	"strings"

	types "ws-server/types"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host      = "localhost"
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

func CreateRoom(UsersID []int) bool {
	db := dbOpen()
	defer db.Close()
	insert, err := db.Prepare("insert into Rooms (UsersID) value (?)")
	checkError(err)
	defer insert.Close()

	sort.Ints(UsersID)
	_, errExec := insert.Exec(int2string(UsersID))
	if errExec != nil {
		checkError(errExec)
		return false
	} else {
		checkError(errExec)
		return true
	}
}

func SelectRoom(UsersID []int) types.Room {
	db := dbOpen()
	defer db.Close()
	selected, err := db.Query("SELECT * FROM Rooms WHERE UsersID=?", int2string(UsersID))
	checkError(err)
	defer selected.Close()

	var room types.Room
	for selected.Next() {
		err = selected.Scan(&room.ID, &room.UsersID)
		if err != nil {
			panic(err)
			return room
		}
	}
	return room
}

func QueryOldMessages(RoomID int) []types.Message {
	var msg types.Message
	var messages []types.Message
	var str, firstname, lastname string
	db := dbOpen()
	defer db.Close()

	stmt := `select RoomID, SenderID, MsgContent, UsersID, firstname, lastname 
				from Message 
				inner join Rooms  
				ON Message.RoomID = Rooms.id
				inner join Users
				ON Message.SenderID = Users.ID
				WHERE RoomID = ?
				ORDER BY Message.id DESC LIMIT 10;
	`
	rows, err := db.Query(stmt, RoomID)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&RoomID, &msg.ID, &msg.Context, &str, &firstname, &lastname)
		msg.Receiver = string2int(str)
		msg.Name = firstname + " " + lastname
		msg.Event = "message"

		checkError(err)
		messages = append(messages, msg)
	}
	err = rows.Err()
	checkError(err)
	return messages
}

func SaveMessages(RoomID int, msg types.Message) {
	db := dbOpen()
	defer db.Close()

	stmt := "insert into Message (RoomID,SenderID,MsgContent) value ( ?, ?, ?)"
	insert, err := db.Prepare(stmt)
	checkError(err)
	defer insert.Close()

	_, errExec := insert.Exec(RoomID, msg.ID, msg.Context)
	checkError(errExec)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func int2string(arr []int) string {
	sort.Ints(arr)
	return strings.Trim(strings.Replace(fmt.Sprint(arr), " ", ",", -1), "[]")
}

func string2int(str string) []int {
	strs := strings.Split(str, ",")
	ary := make([]int, len(strs))
	for i := range ary {
		ary[i], _ = strconv.Atoi(strs[i])
	}
	return ary
}

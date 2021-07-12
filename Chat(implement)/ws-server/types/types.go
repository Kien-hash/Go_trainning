package types

import (
	"github.com/gorilla/websocket"
)

type Message struct {
	Event    string
	ID       int
	Name     string
	Context  string
	Receiver []int
	Users    []User
}

type User struct {
	ID      int
	Name    string
	Address *websocket.Conn
}

type Room struct {
	ID          int
	UsersID     string
	MsgTable    string
}

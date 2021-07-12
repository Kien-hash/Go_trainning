package msghandle

import (
	"log"
	db "ws-server/database"
	types "ws-server/types"

	"github.com/gorilla/websocket"
)

func SendMsg(m map[types.User]bool, msg types.Message) {
	for client := range m {
		if contains(msg.Receiver, client.ID) {
			err := client.Address.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Address.Close()
				delete(m, client)
			}
		}
	}
}

func SendNotice(m map[types.User]bool, msg types.Message) {
	for client := range m {
		err := client.Address.WriteJSON(msg)
		if err != nil {
			log.Printf("error: %v", err)
			client.Address.Close()
			delete(m, client)
		}
	}
}

func InitialUser(m map[types.User]bool, msg types.Message, ws *websocket.Conn) {
	for client := range m {
		if client.ID == msg.ID {
			ws.Close()
			delete(m, client)
		}
	}
	var user types.User
	user.ID = msg.ID
	user.Name = msg.Name
	user.Address = ws
	m[user] = true
}

func InitialSingleRoom(m map[types.User]bool, msg types.Message) types.Room {
	// Truy vấn phòng trong database để tìm tin nhắn cũ
	room := db.SelectRoom(msg.Receiver)
	if room.ID == 0 {
		db.CreateRoom(msg.Receiver)
		room = db.SelectRoom(msg.Receiver)
	}

	var notice types.Message
	for i := range msg.Receiver {
		if msg.Receiver[i] == msg.ID {

		} else {
			notice.Receiver = append(notice.Receiver, msg.Receiver[i])
		}
	}
	notice.ID = msg.ID
	notice.Event = "has-message"
	notice.Name = msg.Name
	notice.Context = room.UsersID
	// log.Println(notice)
	for client := range m {
		if contains(notice.Receiver, client.ID) {
			err := client.Address.WriteJSON(notice)
			// log.Println(notice)
			if err != nil {
				log.Printf("error: %v", err)
				client.Address.Close()
				delete(m, client)
			}
		}
	}
	return room
}

func GetOldMsg(RoomID int ) []types.Message {
	return db.QueryOldMessages(RoomID)
}

func UserOnline(m map[types.User]bool) types.Message {
	var notice types.Message
	notice.Event = "user-online"

	notice.Users = make([]types.User, 0, len(m))
	for k := range m {
		N := k
		N.Address = nil
		notice.Users = append(notice.Users, N)
	}
	return notice
}

func SaveMsg(msg types.Message) {
	room := db.SelectRoom(msg.Receiver)
	db.SaveMessages(room.ID, msg)
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

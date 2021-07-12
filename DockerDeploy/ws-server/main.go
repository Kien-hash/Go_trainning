package main

import (
	"log"
	"net/http"
	msghandle "ws-server/msghandle"
	types "ws-server/types"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var clients = make(map[types.User]bool) // Mảng tham chiếu
var msgChan = make(chan types.Message)
var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	route := mux.NewRouter().StrictSlash(true)
	route.HandleFunc("/ws", handleWebsocket)

	go handleMsgChannel()
	log.Println("http server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", route))
}

func handleMsgChannel() {
	for {
		msg := <-msgChan
		log.Printf("Send to client: %v", msg)

		if msg.Event != "user-online" {
			msghandle.SendMsg(clients, msg)
			continue
		} else {
			msghandle.SendNotice(clients, msg)
		}
	}
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	checkErr(err)
	defer ws.Close()

	for {
		var msg types.Message
		// var room types.Room
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			for client := range clients {
				if client.Address == ws {
					ws.Close()
					delete(clients, client)
				}
			}
			msgChan <- msghandle.UserOnline(clients)
			break
		}
		log.Printf("Receive from client: %v", msg)

		if msg.Event == "initial" {

			if msg.Name == "" || msg.ID < 1 {
				log.Println("wrong username")
				continue
			}
			msghandle.InitialUser(clients, msg, ws)
			msgChan <- msghandle.UserOnline(clients)

		} else if msg.Event == "initial-single-room" {

			if len(msg.Receiver) == 0 {

			} else {
				room := msghandle.InitialSingleRoom(clients, msg)
				msgs := msghandle.GetOldMsg(room.ID)
				for i := len(msgs) - 1; i >= 0; i-- {
					msgChan <- msgs[i]
				}

			}

		} else if msg.Event == "message" {

			msghandle.SaveMsg(msg)
			msgChan <- msg

		}

	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

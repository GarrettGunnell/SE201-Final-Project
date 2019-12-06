package main
/*
import(
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

string chat_message = ''

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan *string)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	}
}

// func root handler(w http.ResponseWriter, r *http.Request){}?

func writer(coord *string){
	broadcast <= coord
}

// func chatHandler(w http.ResponseWriter, r *http.Request){}?

func wsHandler(w http.ResponseWriter, r *http.Request){
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		log.Fatal(err)
	}

	clients[ws]= true
}

func echo(){
	for{
		val := <-broadcast
		message := chat_message

		for client := range clients{
			err := client.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil{
				log.Printf("Websocket error", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
*/

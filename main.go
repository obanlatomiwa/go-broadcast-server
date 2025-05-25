package main

import (
	broadCastServer "github.com/obanlatomiwa/go-broadcast-server/websocket"
	"net/http"
)

func main() {
	//cmd.Execute()
	InitiateBroadCast()
}

func InitiateBroadCast() {
	server := broadCastServer.NewBroadCastServer(":3000")
	go server.StartBroadCast()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		broadCastServer.StartWebsocketConnection(server, w, r)
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return
	}
}

package websocket

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const (
	// max message size allowed from connection peer
	maxMessageSize = 512

	// time allowed to read the next pong message from the peer
	pongWait = 60 * time.Second

	// sends pings to peer. Must be less than pongWait
	pingPeriod = (pongWait * 9) / 10

	writeWait = 10 * time.Second
)

type Client struct {
	id              string
	broadCastServer *BroadCastServer
	conn            *websocket.Conn
	send            chan []byte
}

var connectionUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func StartWebsocketConnection(server *BroadCastServer, writer http.ResponseWriter, request *http.Request) {
	// upgrade the connection from http protocol to websocket protocol
	conn, err := connectionUpgrade.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// create the client
	id := uuid.New()

	client := &Client{
		id:              id.String(),
		broadCastServer: server,
		conn:            conn,
		send:            make(chan []byte, 256),
	}

	// listen to the broadcast server
	client.broadCastServer.register <- client

	go client.readFromBroadCastToWebsocket()
	go client.readFromWebsocketConnection()
}

func (c *Client) readFromWebsocketConnection() {
	defer func() {
		c.broadCastServer.unregister <- c
		err := c.conn.Close()
		if err != nil {
			return
		}
	}()

	c.conn.SetReadLimit(maxMessageSize)

	err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		return
	}

	c.conn.SetPongHandler(func(string) error {
		err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			return err
		}
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		log.Printf("Message from %v: %v", c.id, string(message))
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &Message{
			clientId: c.id,
			data:     string(message),
		}

		c.broadCastServer.broadcast <- msg
	}
}

func (c *Client) readFromBroadCastToWebsocket() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		err := c.conn.Close()
		if err != nil {
			return
		}
	}()

	for {
		select {
		case message, ok := <-c.send:
			err := c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				return
			}

			if !ok {
				err := c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					return
				}
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, err = w.Write(message)
			if err != nil {
				return
			}

			// for performance, Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				_, err := w.Write(message)
				if err != nil {
					return
				}
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			err := c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				return
			}
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}

	}

}

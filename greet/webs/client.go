package webs

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID         string
	Socket     *websocket.Conn
	Send       chan string
	Register   chan *Client
	Unregister chan *Client
}

var clients = make(map[*websocket.Conn]*Client)

var broadcast = make(chan string)

func init() {
	go func() {
		for {
			msg := <-broadcast
			for _, client := range clients {
				client.Send <- msg
			}
		}
	}()
}

func WebSocketHandler(writer http.ResponseWriter, request *http.Request) {
	upgrader = websocket.Upgrader{}
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	client := &Client{
		Socket:     conn,
		Send:       make(chan string, 256),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}

	clients[conn] = client

	go client.Read()
	go client.Write()

	client.Register <- client
}

func (c *Client) Read() {
	defer func() {
		c.Unregister <- c
		close(c.Send)
	}()
	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			return
		}
		// 这里可以处理接收到的消息，例如广播给其他客户端
		broadcast <- string(message)
	}
}

func (c *Client) Write() {
	defer func() {
		c.Socket.Close()
	}()
	for {
		select {
		case message := <-c.Send:
			err := c.Socket.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("Write error:", err)
				return
			}
		}
	}
}

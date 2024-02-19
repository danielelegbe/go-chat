package ws

import (
	"log"

	"github.com/danielelegbe/go-chat/database"
	"github.com/danielelegbe/go-chat/models"
	"github.com/gorilla/websocket"
)

type ClientMessage struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

type Client struct {
	conn    *websocket.Conn
	receive chan *ClientMessage
	room    *Room
}

func (c *Client) read() {
	defer c.conn.Close()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var newMessage models.Message
		newMessage.Content = string(msg)

		result := database.DB.Create(&newMessage)

		if result.Error != nil {
			log.Println(result.Error)
		}

		c.room.broadcast <- &ClientMessage{
			ID:      newMessage.ID,
			Content: newMessage.Content,
		}
	}
}

func (c *Client) write() {
	defer c.conn.Close()

	for msg := range c.receive {

		err := c.conn.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}

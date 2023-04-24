package streamer

import (
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type ReceiveData struct {
	Id      uuid.UUID
	Payload []byte
}

type Client struct {
	Id       uuid.UUID
	conn     *websocket.Conn
	receiver chan ReceiveData
	sender   chan []byte
	closer   chan bool
	Active   bool
}

func newClient(conn *websocket.Conn, receiver chan ReceiveData) *Client {
	return &Client{
		Id:       uuid.Must(uuid.NewV4()),
		conn:     conn,
		receiver: receiver,
		sender:   make(chan []byte),
		closer:   make(chan bool),
		Active:   true,
	}
}
func (c *Client) listen() {
	for {
		messageType, message, err := c.conn.ReadMessage()
		if err != nil {
			c.closer <- true
			return
		}
		if messageType != websocket.TextMessage {
			continue
		}
		// fmt.Printf("message: %s\n", message)

		c.receiver <- ReceiveData{
			Id:      c.Id,
			Payload: message,
		}
	}
}

func (c *Client) send() {
	for {
		message, ok := <-c.sender

		if !ok {
			c.closer <- true
			return
		}

		err := c.conn.WriteMessage(websocket.TextMessage, message)

		if err != nil {
			c.closer <- true
			return
		}
		// fmt.Printf("sent: %s\n", message)
	}
}

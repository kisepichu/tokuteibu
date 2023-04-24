package streamer

import (
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

type Streamer struct {
	Clients  map[uuid.UUID]*Client
	receiver chan ReceiveData
}

func NewStreamer() *Streamer {
	return &Streamer{
		Clients:  make(map[uuid.UUID]*Client),
		receiver: make(chan ReceiveData),
	}
}

type Payload struct {
	Type string                 `json:"type,omitempty"`
	Body   map[string]interface{} `json:"body,omitempty"`
}

func (s *Streamer) Listen(handlerWebSocket func(s *Streamer, data ReceiveData) error) {
	for {
		data := <-s.receiver

		go func() {
			err := handlerWebSocket(s, data)
			if err != nil {
				log.Print("error: ", err)
			}
		}()
	}
}

func (s *Streamer) Send(message []byte, cond func(c *Client) bool) error {
	for _, c := range s.Clients {
		if cond(c) {
			c.sender <- message
		}
	}
	return nil
}

func (s *Streamer) SendTo(id uuid.UUID, message []byte) error {
	c, ok := s.Clients[id]
	if !ok {
		return fmt.Errorf("client not found")
	}
	c.sender <- message
	return nil
}

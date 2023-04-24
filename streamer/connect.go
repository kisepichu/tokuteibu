package streamer

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *Streamer) ConnectWS(c echo.Context, whenClosed func(c *Client)) error {
	connection, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	defer connection.Close()

	client := newClient(connection, s.receiver) // receiver is shared by streamer and all clients

	s.Clients[client.Id] = client
	go client.listen()
	go client.send()

	<-client.closer

	whenClosed(client)

	return c.NoContent(http.StatusOK)
}

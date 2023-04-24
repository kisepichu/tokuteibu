package main

import (
	"net/http"

	"tokuteibu/handlers"
	"tokuteibu/streamer"

	"github.com/labstack/echo/v4"
)

func main(){
	s := streamer.NewStreamer()
	e:= echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "front: mada")
	})

	api := e.Group("/api")
	{
		api.GET("/ping", func(c echo.Context) error {
			return c.String(http.StatusOK, "pong")
		})
		api.GET("/ws", func(c echo.Context) error {
			s.ConnectWS(c, func(c *streamer.Client) {})
			return nil
		})
	}
	go s.Listen(handlers.HandlerWebSocket)

	e.Logger.Panic(e.Start(":3939"))
}

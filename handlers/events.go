package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Event struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Owner string `json:"owner"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func HandlerGetEvents(c echo.Context) error {
	var events []Event
	return c.JSON(http.StatusOK, events)
}

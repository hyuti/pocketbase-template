package controller

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func RegisterHealthRoute(handler *Router) {
	handler.Get("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
}

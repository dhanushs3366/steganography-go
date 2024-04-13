package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Router struct {
	*echo.Echo
}

func (r *Router) GetEncode(c echo.Context) error {
	return c.String(http.StatusOK, "Encoding")
}

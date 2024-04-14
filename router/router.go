package router

import (
	pages "steganography/views/pages"

	"github.com/labstack/echo/v4"
)

type Router struct {
	*echo.Echo
}

func (r *Router) GetEncode(c echo.Context) error {
	return pages.GetEncode().Render(c.Request().Context(), c.Response().Writer)
}

func (r *Router) PostEncode(c echo.Context) error {
	return c.String(200, "Encodedddd")
}
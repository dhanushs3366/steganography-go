package main

import (
	router "steganography/router"
	views "steganography/views"

	"github.com/labstack/echo/v4"
)

func main() {

	r := &router.Router{Echo: echo.New()}

	r.GET("/", func(c echo.Context) error {
		return views.Boilerplate().Render(c.Request().Context(), c.Response().Writer)
	})
	r.GET("/encode", r.GetEncode)

	r.Logger.Fatal(r.Start(":8888"))
}

package main

import (
	"steganography/consts"
	"steganography/router"
	"steganography/views"

	"github.com/labstack/echo/v4"
)

func main() {

	r := &router.Router{Echo: echo.New()}

	r.Static("/output", "output")
	r.Static("/static", "static")
	r.GET("/", func(c echo.Context) error {
		consts.CLIENT_STATE = "HOME"
		return views.Boilerplate().Render(c.Request().Context(), c.Response().Writer)
	})

	r.GET("/encode", r.GetEncode)
	r.POST("/encode", r.PostEncode)
	r.GET("/encode/:filename", func(c echo.Context) error {

		filename := c.Param("filename")

		return c.Attachment("output/"+filename, filename)
	})

	r.GET("/decode", r.GetDecode)
	r.POST("/decode", r.PostDecode)
	r.Logger.Fatal(r.Start(":8080"))

}

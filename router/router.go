package router

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"path/filepath"
	"steganography/stegano"
	pages "steganography/views/pages"
	"strings"

	"github.com/labstack/echo/v4"
)

const outputDir = "output"

type Router struct {
	*echo.Echo
}

func (r *Router) GetEncode(c echo.Context) error {
	return pages.GetEncode().Render(c.Request().Context(), c.Response().Writer)
}

func (r *Router) PostEncode(c echo.Context) error {
	file, err := c.FormFile("encode-image")
	if err != nil {
		fmt.Println(err)
		return err
	}
	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer src.Close()

	// Check if the uploaded file is PNG
	ext := filepath.Ext(file.Filename)
	isPNG := strings.EqualFold(ext, ".png")

	var img image.Image
	fmt.Println("PNG sec initiation")
	if isPNG {
		// Decode the PNG image
		img, err = png.Decode(src)
		if err != nil {

			fmt.Println(err)
			return err
		}
	} else {
		// Assume other image types are already JPEG or similar
		img, _, err = image.Decode(src)
		if err != nil {

			fmt.Println(err)
			return err
		}
	}

	// Destination
	dstName := "encode"
	if isPNG {
		dstName = strings.TrimSuffix(dstName, ext) + ".jpg"
	}
	dst, err := os.Create(outputDir + "/" + dstName)
	if err != nil {

		fmt.Println(err)
		return err
	}
	defer dst.Close()

	// Encode the image as JPEG
	err = jpeg.Encode(dst, img, nil)
	if err != nil {

		fmt.Println(err)
		return err
	}

	newPath := path.Join(outputDir, dstName)

	err = os.Rename(dst.Name(), newPath)
	if err != nil {

		fmt.Println(err)
		return err
	}

	err = stegano.Encode(newPath, "Hello World")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return pages.RenderEncode("encoded.png").Render(c.Request().Context(), c.Response().Writer)
}

func (r *Router) GetEncodedImage(fileName string, c echo.Context) error {
	filename := c.Param(fileName)

	return c.Attachment(filename, filename)
}

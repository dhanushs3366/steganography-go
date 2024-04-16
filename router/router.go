package router

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
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

	img, err := png.Decode(src)
	if err != nil {
		fmt.Println(err)
		return err
	}

	dstName := "encode.png"
	dst, err := os.Create(outputDir + "/" + dstName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer dst.Close()
	if err = png.Encode(dst, img); err != nil {
		fmt.Println(err)
		return err
	}
	return pages.RenderEncode(dstName).Render(c.Request().Context(), c.Response().Writer)
}

func (r *Router) PngToJpg(src io.Reader) (image.Image, error) {
	img, err := png.Decode(src)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return img, nil
}

func (r *Router) GetEncodedImage(fileName string, c echo.Context) error {
	filename := c.Param(fileName)

	return c.Attachment(filename, filename)
}

func (r *Router) GetDecode(c echo.Context) error {
	return pages.GetDecode().Render(c.Request().Context(), c.Response().Writer)
}

func (r *Router) PostDecode(c echo.Context) error {
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

	ext := filepath.Ext(file.Filename)
	isPNG := strings.EqualFold(ext, ".png")

	var img image.Image
	if isPNG {
		img, err = r.PngToJpg(src)
		if err != nil {
			return err
		}
	} else {
		img, _, err = image.Decode(src)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	dstName := "decode"
	if isPNG {
		dstName = strings.TrimSuffix(dstName, ext) + ".jpg"
	}
	dst, err := os.Create(outputDir + "/" + dstName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer dst.Close()

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

	text, err := stegano.Decode(newPath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	println(text)
	return pages.RenderDecode(text).Render(c.Request().Context(), c.Response().Writer)

}

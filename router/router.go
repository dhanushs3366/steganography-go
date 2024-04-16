package router

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"net/http"
	"net/url"
	"os"
	pages "steganography/views/pages"

	"github.com/labstack/echo/v4"
)

const outputDir = "output"
const pyScriptURL = "http://localhost:3000"

type Router struct {
	*echo.Echo
}

func (r *Router) GetEncode(c echo.Context) error {
	return pages.GetEncode().Render(c.Request().Context(), c.Response().Writer)
}

func (r *Router) PostEncode(c echo.Context) error {

	file, err := c.FormFile("encode-image")
	text := c.FormValue("encode-text")
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

	// image is saved locally by python server just call the request dont need to save the mime type received from response from python server just need a confirmation to proceed

	formData := url.Values{}
	formData.Add("text", text) // Use "text" as the key
	res, err := http.PostForm(pyScriptURL+"/encode", formData)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
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
	// Retrieve the uploaded image file
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

	// Decode the image
	img, _, err := image.Decode(src)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Create a destination file with PNG extension
	dstName := "decode.png"
	dst, err := os.Create(outputDir + "/" + dstName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer dst.Close()

	// Encode the image as PNG and save to the destination file
	if err := png.Encode(dst, img); err != nil {
		fmt.Println(err)
		return err
	}

	// Send a request to the Python server
	formData := url.Values{}
	formData.Add("image_path", "../output/encoded.png")
	res, err := http.PostForm(pyScriptURL+"/decode", formData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	// Read and print the response body
	content, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	decoded_txt := string(content)

	return pages.RenderDecode(decoded_txt).Render(c.Request().Context(), c.Response().Writer)
}

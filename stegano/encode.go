package stegano

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"os"
)

// Encode embeds a text message into an image using LSB steganography
func Encode(imagePath string, message string) error {
	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	bounds := img.Bounds()
	maxChars := (bounds.Dx() * bounds.Dy() * 3) / 8
	if len(message) > maxChars {
		return fmt.Errorf("message is too long to embed in the image")
	}

	// Create a new RGBA image to modify pixels
	rgba := image.NewRGBA(bounds)

	// Embed the message in the pixels
	charIndex := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			// Embed character in the least significant bits of each color channel
			if charIndex < len(message) {
				char := message[charIndex]
				r = embedChar(r, char)
				charIndex++
			}
			if charIndex < len(message) {
				char := message[charIndex]
				g = embedChar(g, char)
				charIndex++
			}
			if charIndex < len(message) {
				char := message[charIndex]
				b = embedChar(b, char)
				charIndex++
			}

			// Set the modified pixel color
			rgba.Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 255})
		}
	}

	// Create a new image file with the embedded message
	outputFile, err := os.Create("output/encoded.png")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Encode the modified image to PNG format
	err = png.Encode(outputFile, rgba)
	if err != nil {
		return err
	}

	return nil
}

func embedChar(color uint32, char byte) uint32 {
	// Embed the character in the least significant bits of the color channel
	return (color & 0xfffffff8) | uint32(char)
}

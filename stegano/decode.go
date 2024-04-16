package stegano

import (
	"fmt"
	"image"
	"os"
)

// Decode extracts a text message from an image using LSB steganography
func Decode(imagePath string) (string, error) {
	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	bounds := img.Bounds()

	// Initialize variables to store the extracted message and character index
	var message string
	var charIndex int

	// Iterate through the pixels to extract the message
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			// Extract character from the least significant bits of each color channel
			if charIndex < 8 {
				char := byte(extractChar(r))
				message += string(char)
				charIndex++
			}
			if charIndex < 8 {
				char := byte(extractChar(g))
				message += string(char)
				charIndex++
			}
			if charIndex < 8 {
				char := byte(extractChar(b))
				message += string(char)
				charIndex++
			}

			// Break the loop if the end of message delimiter is found
			if len(message) >= 8 && message[len(message)-8:] == "########" {
				message = message[:len(message)-8]
				return message, nil
			}
		}
	}
	return "", fmt.Errorf("end of message delimiter not found")
}

func extractChar(color uint32) byte {
	// Extract the character from the least significant bits of the color channel
	return byte(color & 0x00000007)
}

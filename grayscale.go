package mymodule

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func Grayscale(inputPath, outputPath string) error {
	// Open the original image
	reader, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("error opening the image: %v", err)
	}
	defer reader.Close()

	// Decode the image
	img, _, err := image.Decode(reader)
	if err != nil {
		return fmt.Errorf("error decoding the image: %v", err)
	}

	// Get image bounds
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Create a new grayscale image
	grayImg := image.NewGray(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Get the color at pixel (x, y)
			oldColor := img.At(x, y)
			r, g, b, _ := oldColor.RGBA()

			// Convert to gray using the formula
			gray := uint8((0.3*float64(r) + 0.59*float64(g) + 0.11*float64(b)) / 256.0)

			// Set the gray color
			grayColor := color.Gray{Y: gray}
			grayImg.Set(x, y, grayColor)
		}
	}

	// Save the grayscale image
	grayFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating the grayscale image file: %v", err)
	}
	defer grayFile.Close()
	png.Encode(grayFile, grayImg)

	fmt.Println("Grayscale image saved.")
	return nil
}

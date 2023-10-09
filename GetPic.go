package mymodule

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// DownloadImage downloads an image from the specified URL and saves it to the specified file.
func GetImage(imageURL, fileName string) error {
	// Create an HTTP GET request
	response, err := http.Get(imageURL)
	if err != nil {
		return fmt.Errorf("error making the request: %v", err)
	}
	defer response.Body.Close()

	// Check if the response status code is OK (200)
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("error: status code %d", response.StatusCode)
	}

	// Create a new file to save the image
	outputFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating the file: %v", err)
	}
	defer outputFile.Close()

	// Copy the HTTP response body to the file
	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		return fmt.Errorf("error saving the image: %v", err)
	}

	fmt.Println("Image downloaded and saved as", fileName)
	return nil
}

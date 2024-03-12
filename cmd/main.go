package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("menu.txt") // Update this to your file path
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the contents of the file into a buffer
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the byte slice to a string
	text := string(content)

	// Find and replace the phrase
	text = strings.Replace(text, " - ", ":", -1)

	// Output the modified content; you might want to write this back to a file
	fmt.Println(text)

	// To write the modified content back to a file, uncomment the following lines:
	// err = os.WriteFile("output_filename.txt", []byte(text), 0644) // Specify your new output file name
	// if err != nil {
	//     log.Fatal(err)
	// }
}
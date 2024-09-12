package main

import (
	"fmt"
	"os"
)

func readFile(path string) (string, error) {
	bytes, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Missing arguments")
		return
	}

	text, readingErr := readFile(args[0])

	if readingErr != nil {
		fmt.Printf("Failed to read the file due to %s", readingErr.Error())
		return
	}

	fmt.Println(text)
}

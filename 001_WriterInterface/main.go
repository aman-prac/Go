package main

import (
	"log"
	"os"
)

func main() {

	f, err := os.Create("Output.txt")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer f.Close()

	s := []byte("Hello, World! This is a test file.\n and no next line ")

	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
}

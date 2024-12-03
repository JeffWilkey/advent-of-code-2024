package utils

import (
	"log"
	"os"
)

func ReadInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	return string(data)
}

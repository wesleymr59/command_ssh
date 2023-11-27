package utils

import (
	"log"
	"os"
	"strings"
)

var entries []string

func ReadFile() []string {
	contentBytes, err := os.ReadFile("cmd.txt")
	if err != nil {
		log.Println("Error reading the file:", err)

	}
	// Convert the byte slice to a string
	content := string(contentBytes)
	entries = strings.SplitN(content, ";", 3)
	return entries
}

package main

import (
	"os"
	"strings"
)

// Helper: Checks if all characters are within printable ASCII range (32-126).
func IsPrintable(text string) bool {
	text = strings.ReplaceAll(text, "\r\n", "")
	for _, char := range text {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}

// Helper: Checks if a file exists.
func fileNotExists(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

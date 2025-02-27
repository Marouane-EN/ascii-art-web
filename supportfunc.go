package main

import "strings"

// IsPrintable checks if all characters in the input are within printable ASCII range (32-126)
func IsPrintable(text string) bool {
	text = strings.ReplaceAll(text, "\r\n", "")
	for _, char := range text {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}

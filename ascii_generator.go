package main

import (
	"fmt"
	"os"
	"strings"
)

// ConvertToAscii converts input text into ASCII art using a chosen banner.
func ConvertToAscii(text []string, banner string) string {
	// Load banner data
	filePath := fmt.Sprintf("banners/%s.txt", banner)
	bannerData, _ := os.ReadFile(filePath)

	// Handle different banner types
	var lines []string
	if banner == "thinkertoy" {
		lines = strings.Split(string(bannerData), "\r\n")
	} else {
		lines = strings.Split(string(bannerData), "\n")
	}

	// Parse the banner to map ASCII characters
	charMap := parseBanner(lines)

	var outputLines [8]string
	var result string

	// Convert input text to ASCII
	for index, word := range text {
		for _, char := range word {
			if asciiRep, exists := charMap[char]; exists {
				for i := range outputLines {
					outputLines[i] += asciiRep[i]
				}
			}
		}

		if word == "" {
			result += "\n"
		} else if len(text) > 1 && index < len(text)-1 {
			result += strings.Join(outputLines[:], "\n") + "\n"
			outputLines = [8]string{}
		}
	}

	if outputLines != [8]string{} {
		result += strings.Join(outputLines[:], "\n") + "\n"
	}
	return result
}

// parseBanner converts banner lines into a map of ASCII art.
func parseBanner(lines []string) map[rune][]string {
	lines = lines[1:] // Skip first empty line
	charMap := make(map[rune][]string)
	asciiCode := 32  // Start at ASCII code 32 (space)

	for i := range lines {
		if lines[i] == "" {
			charMap[rune(asciiCode)] = lines[i-8 : i]
			asciiCode++
		}
	}
	return charMap
}

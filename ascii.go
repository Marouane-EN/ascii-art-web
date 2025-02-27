package main

import (
	"fmt"
	"os"
	"strings"
)

// ConvertToAscii converts input text into ASCII art using the specified banner.
// Banner files should be placed in the "banners/" directory.
func ConvertToAscii(text []string, banner string) string {
	// Build the banner file path (e.g., banners/standard.txt)
	filePath := fmt.Sprintf("banners/%s.txt", banner)
	bannerData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Sprintf("Error: Banner file %s not found", filePath)
	}
	var lines []string
	// Split banner file into lines
	if banner == "thinkertoy" {
		lines = strings.Split(string(bannerData), "\r\n")
	} else {
		lines = strings.Split(string(bannerData), "\n")
	}
	// Build a character map from the banner file (assuming each character is represented by 8 lines)
	charMap, err := parseBanner(lines)
	if err != nil {
		return "Error parsing banner file."
	}

	// Prepare an array for the ASCII art output (8 lines per character)
	outputLines := [8]string{}

	result := ""

	for index, word := range text {
		for _, char := range word {
			asciiRep, exists := charMap[char]
			if !exists {
				continue // Skip characters not in the banner map
			}
			for i := range outputLines {
				outputLines[i] += asciiRep[i]
			}
		}
		if word == "" {
			result += "\n"
			continue
		}
		if len(text) > 1 && index < len(text)-1 {
			result += strings.Join(outputLines[:], "\n") + "\n"
			outputLines = [8]string{}
		}
	}

	if outputLines != [8]string{} {
		result += strings.Join(outputLines[:], "\n") + "\n"
	}
	return result
}

// parseBanner converts banner file lines into a map of runes to their ASCII art representation.
func parseBanner(lines []string) (map[rune][]string, error) {
	lines = lines[1:]
	charMap := make(map[rune][]string)
	asciiCode := 32 // Start at ASCII code 32 (space)
	// Each character block is 8 lines
	for i := range lines {
		if lines[i] == "" { // empty string is a separator between characters
			charMap[rune(asciiCode)] = lines[i-8 : i]
			asciiCode++
		}
	}

	// Ensure we have 95 characters mapped (ASCII 32 to 126)
	if len(charMap) != 95 {
		return charMap, fmt.Errorf("banner file is missing characters")
	}

	return charMap, nil
}

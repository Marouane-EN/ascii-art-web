package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run main.go")
		return
	}

	// Serve static files (like CSS)
	http.HandleFunc("/static/", StaticHandler)
	// Route for home page (GET)
	http.HandleFunc("/", homeHandler)
	// Route for ASCII art generation (POST)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	http.HandleFunc("/ascii-art1", asciiArtHandler)
	port := ":8080"
	fmt.Printf("🚀 Server running at http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("❌ Error starting server:", err)
	}
}

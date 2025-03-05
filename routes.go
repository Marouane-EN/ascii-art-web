package main

import (
	"html/template"
	"net/http"
	"strings"
)

// PageData holds the text input and ASCII art output.
type PageData struct {
	AsciiArt     string
	ErrorMessage string
	Code         int
}

// homeHandler serves the home page.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendErrorPage(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}
	if r.URL.Path != "/" {
		sendErrorPage(w, http.StatusNotFound, "Oops! Page not found")
		return
	}
	renderResultPage(w, &PageData{}, "home.html")
}

// asciiArtHandler processes the form submission, generates ASCII art, and displays the output.
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendErrorPage(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	inputText := r.FormValue("text")
	banner := r.FormValue("banner")
	outputPage := "output.html"
	if r.URL.Path == "/ascii-art" {
		outputPage = "home.html"
	}

	// Input validation
	if inputText == "" {
		sendErrorPage(w, http.StatusBadRequest, "Please enter some text!", outputPage)
		return
	}

	if !IsPrintable(inputText) {
		sendErrorPage(w, http.StatusBadRequest, "Invalid character detected!", outputPage)
		return
	}

	if fileExists("banners/" + banner + ".txt") {
		sendErrorPage(w, http.StatusNotFound, "Banner not found!")
		return
	}

	// Character limit check
	textLen := len(strings.ReplaceAll(inputText, "\r\n", "\n"))
	if textLen > 1000 {
		sendErrorPage(w, http.StatusBadRequest, "Character limit exceeded (max 1000).", outputPage)
		return
	}

	// Process ASCII Art
	textLines := strings.Split(inputText, "\r\n")
	asciiArt := ConvertToAscii(textLines, banner)
	
	outputPage = "output.html"
	
	renderResultPage(w, &PageData{AsciiArt: asciiArt}, outputPage)
}

// StaticHandler serves static files (e.g., CSS, images).
func StaticHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendErrorPage(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	url := r.URL.Path[1:] // Remove leading slash
	if fileExists(url) {
		sendErrorPage(w, http.StatusNotFound, "Static file not found")
		return
	}

	http.ServeFile(w, r, url)
}

// Helper: Renders the given HTML template with provided PageData.
func renderResultPage(w http.ResponseWriter, data *PageData, templateName string) {
	tmpl, err := template.ParseFiles("templates/" + templateName)
	if err != nil {
		sendErrorPage(w, http.StatusInternalServerError, "Oops! Something went wrong.")
		return
	}
	tmpl.Execute(w, data)
}

// Helper: Sends an error page with appropriate HTTP status code.
func sendErrorPage(w http.ResponseWriter, code int, message string, page ...string) {
	errorPage := "error.html"
	if len(page) > 0 {
		errorPage = page[0]
	}
	w.WriteHeader(code)

	renderResultPage(w, &PageData{
		Code:         code,
		ErrorMessage: message,
	}, errorPage)

}

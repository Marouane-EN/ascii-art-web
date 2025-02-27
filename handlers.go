package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

// PageData holds the text input and ASCII art result.
type PageData struct {
	InputText    string
	AsciiArt     string
	Banner       string
	ErrorMessage string
}

// homeHandler serves the main page with the text area and banner selection.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}
	renderResultPage(w, &PageData{}, "index.html")
}

// asciiArtHandler processes the form submission, generates ASCII art, and displays the result.
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Allow both GET and POST on this route.
	// On GET, simply display the result page with an empty form.
	// fmt.Printf("%#v\n", r.URL)
	// fmt.Printf("%#v\n", r.URL.Path)

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}
	// Read form values
	inputText := r.FormValue("text")
	fmt.Printf("%#v\n", inputText)

	banner := r.FormValue("banner")
	fmt.Printf("%#v\n", banner)

	// Validate input: if empty, set an error message.
	if inputText == "" && r.URL.Path == "/ascii-art" {
		w.WriteHeader(http.StatusBadRequest)
		renderResultPage(w, &PageData{
			ErrorMessage: "Please enter some text!",
		}, "index.html")
		return
	}
	if inputText == "" && r.URL.Path == "/ascii-art1" {
		w.WriteHeader(http.StatusBadRequest)

		renderResultPage(w, &PageData{
			ErrorMessage: "Please enter some text!",
		}, "result.html")
		return
	}
	if !IsPrintable(inputText) && r.URL.Path == "/ascii-art" {
		w.WriteHeader(http.StatusBadRequest)

		renderResultPage(w, &PageData{
			ErrorMessage: "you have invalid character",
		}, "index.html")
		return
	}
	if !IsPrintable(inputText) && r.URL.Path == "/ascii-art1" {
		w.WriteHeader(http.StatusBadRequest)

		renderResultPage(w, &PageData{
			ErrorMessage: "you have invalid character",
			Banner:       banner,
		}, "result.html")
		return
	}

	if banner == "" && r.URL.Path == "/ascii-art" {
		w.WriteHeader(http.StatusBadRequest)

		renderResultPage(w, &PageData{
			ErrorMessage: "choose ascii-art type",
		}, "index.html")
		return
	}
	if banner == "" && r.URL.Path == "/ascii-art1" {
		w.WriteHeader(http.StatusBadRequest)

		renderResultPage(w, &PageData{
			ErrorMessage: "choose ascii-art type",
			Banner:       banner,
		}, "result.html")
		return
	}
	// Replace escaped newlines with actual newlines
	Text := strings.Split(inputText, "\r\n")

	textlen := len(strings.Join(Text, ""))

	if textlen > 1000 && r.URL.Path == "/ascii-art" {
		w.WriteHeader(http.StatusBadRequest)

		renderResultPage(w, &PageData{
			ErrorMessage: "you have exceeded the character limit",
			Banner:       banner,
		}, "index.html")
		return
	}

	if textlen > 1000 && r.URL.Path == "/ascii-art1" {
		w.WriteHeader(http.StatusBadRequest)

		renderResultPage(w, &PageData{
			ErrorMessage: "you have exceeded the character limit",
			Banner:       banner,
		}, "result.html")
		return
	}
	// Generate ASCII art from input text using the selected banner
	asciiArt := ConvertToAscii(Text, banner)

	// Prepare data to pass to the template
	data := PageData{
		InputText: inputText,
		AsciiArt:  asciiArt,
		Banner:    banner,
	}

	renderResultPage(w, &data, "result.html")
}

// renderResultPage loads and executes the result template.
func renderResultPage(w http.ResponseWriter, data *PageData, htmltemplate string) {
	tmpl, err := template.ParseFiles("templates/" + htmltemplate)
	if err != nil {
		http.Error(w, "Error loading result page", http.StatusInternalServerError)
		fmt.Println("Template error:", err)
		return
	}
	tmpl.Execute(w, data)
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", 405)
		return
	}
	url := r.URL.Path
	fmt.Println(url)
	f, err := os.Stat(url[1:])
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "not found", 404)
			return
		}
		http.Error(w, "internal server error", 500)
		return
	}
	if f.IsDir() {
		http.Error(w, "not found", 404)
		return
	}
	http.ServeFile(w, r, url[1:])
}

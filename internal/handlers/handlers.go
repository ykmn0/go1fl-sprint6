package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"go1fl-sprint6/internal/service"
)

// IndexHandler serves the index.html page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "index.html")
}

// UploadHandler handles file uploads and conversion
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	// Get the file from form
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Read the file content
	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	// Convert the content
	result, err := service.Convert(string(content))
	if err != nil {
		http.Error(w, "Error converting content: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate output filename
	ext := filepath.Ext(handler.Filename)
	timestamp := time.Now().UTC().Format("20060102-150405")
	outputFilename := fmt.Sprintf("converted_%s%s", timestamp, ext)

	// Create output file
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		http.Error(w, "Error creating output file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// Write the converted content
	_, err = outputFile.WriteString(result)
	if err != nil {
		http.Error(w, "Error writing output file", http.StatusInternalServerError)
		return
	}

	// Return the conversion result
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(result))
}

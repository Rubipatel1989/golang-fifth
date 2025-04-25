package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"visionApi/utils"
)

func CompareHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get profile image
	profileFile, _, err := r.FormFile("profile")
	if err != nil {
		http.Error(w, "Profile image is required", http.StatusBadRequest)
		return
	}
	defer profileFile.Close()

	// Get ID image
	idFile, _, err := r.FormFile("idcard")
	if err != nil {
		http.Error(w, "ID card image is required", http.StatusBadRequest)
		return
	}
	defer idFile.Close()

	// Save images temporarily
	profilePath := filepath.Join(os.TempDir(), "profile.jpg")
	idPath := filepath.Join(os.TempDir(), "id.jpg")

	profileData, _ := ioutil.ReadAll(profileFile)
	idData, _ := ioutil.ReadAll(idFile)

	os.WriteFile(profilePath, profileData, 0644)
	os.WriteFile(idPath, idData, 0644)

	// Run OCR
	text := utils.ExtractTextFromImage(idPath)

	// Face Match
	similarity := utils.CompareFaces(profilePath, idPath)

	// Respond
	fmt.Fprintf(w, "Extracted Text:\n%s\n\nFace Similarity: %.2f", text, similarity)
}

package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"vision-api-go/utils"
)

func CompareHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	profileFile, _, err := r.FormFile("profile")
	if err != nil {
		http.Error(w, "Missing profile image", http.StatusBadRequest)
		return
	}
	defer profileFile.Close()

	idFile, _, err := r.FormFile("idcard")
	if err != nil {
		http.Error(w, "Missing ID card image", http.StatusBadRequest)
		return
	}
	defer idFile.Close()

	profilePath := filepath.Join("uploads", "profile.jpg")
	idPath := filepath.Join("uploads", "idcard.jpg")

	saveFile(profileFile, profilePath)
	saveFile(idFile, idPath)

	text := utils.ExtractText(idPath)
	similarity := utils.CompareFaces(profilePath, idPath)

	fmt.Fprintf(w, "📝 Extracted Text: \n%s\n\n👤 Face Match Score: %v", text, similarity)
}

func saveFile(src io.Reader, dstPath string) {
	dst, _ := os.Create(dstPath)
	defer dst.Close()
	io.Copy(dst, src)
}

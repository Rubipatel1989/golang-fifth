package utils

import (
	"log"

	"github.com/otiai10/gosseract/v2"
)

func ExtractTextFromImage(imagePath string) string {
	client := gosseract.NewClient()
	defer client.Close()

	err := client.SetImage(imagePath)
	if err != nil {
		log.Println("OCR SetImage error:", err)
		return ""
	}

	text, err := client.Text()
	if err != nil {
		log.Println("OCR Text error:", err)
		return ""
	}

	return text
}

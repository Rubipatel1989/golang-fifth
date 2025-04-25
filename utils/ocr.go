package utils

import (
	"os/exec"
)

func ExtractText(imagePath string) string {
	out, err := exec.Command("tesseract", imagePath, "stdout").Output()
	if err != nil {
		return "OCR failed"
	}
	return string(out)
}

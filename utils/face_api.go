package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const (
	apiKey     = "rk_a_rTbfJzudhZD_KUtu47iyXQMpbRr"
	apiSecret  = "yv4vXMTCs_LWiao3EcMAoHucKs9pSk6n"
	compareURL = "https://api-us.faceplusplus.com/facepp/v3/compare"
)

func CompareFaces(img1Path, img2Path string) float64 {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("api_key", apiKey)
	writer.WriteField("api_secret", apiSecret)
	addFile(writer, "image_file1", img1Path)
	addFile(writer, "image_file2", img2Path)
	writer.Close()

	req, _ := http.NewRequest("POST", compareURL, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if score, ok := result["confidence"].(float64); ok {
		return score
	}
	return 0
}

func addFile(writer *multipart.Writer, fieldName, path string) {
	file, _ := os.Open(path)
	defer file.Close()

	part, _ := writer.CreateFormFile(fieldName, filepath.Base(path))
	io.Copy(part, file)
}

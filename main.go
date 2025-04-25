package main

import (
	"fmt"
	"net/http"

	"visionApi/handlers"
)

func main() {
	http.HandleFunc("/api/compare", handlers.CompareHandler)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net/http"
	"vision-api-go/handlers"
)

func main() {
	http.HandleFunc("/api/compare", handlers.CompareHandler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

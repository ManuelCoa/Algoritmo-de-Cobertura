package main

import (
	"log"
	"net/http"
	"set-cover-service/handlers"
)

func main() {
	http.HandleFunc("/cover", handlers.CoverHandler)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

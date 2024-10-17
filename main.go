package main

import (
	"fmt"
	"log"
	"net/http"

	"asciiart/serv"
)

func main() {


	http.HandleFunc("/", serv.Index)
	http.HandleFunc("/ascii-art", serv.AsciiWeb)
	http.HandleFunc("/ascii-art/export", serv.ExportAsciiArt)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Serving web folder on port 8000")
	http.ListenAndServe(":8000", http.FileServer(http.Dir("web")))
}

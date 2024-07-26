package main

import (
	"log"
	"net/http"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, "templates/index.html")
}

func about(w http.ResponseWriter, r *http.Request)  {
	
}

func main() {

	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/about", about)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

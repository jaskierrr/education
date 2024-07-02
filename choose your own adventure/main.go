package main

import (
	"fmt"
	"net/http"

	"main/handlers"

	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()

	r.HandleFunc("/{key}", handlers.WriteStory)
	r.HandleFunc("/", handlers.Void)

	fmt.Println("Starting the server on port :8080...")
	http.ListenAndServe(":8080", r)
}

package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	handlers "main/handlers"
	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}


func main() {
	dataPath := flag.String("data", "data.yaml", "choose YAML or JSON file with data")
	flag.Parse()

	fmt.Printf("Data from: %q\n",*dataPath)

	r := mux.NewRouter()

	r.HandleFunc("/", hello).Methods("GET")

	r.HandleFunc("/set/{key}/{value}", handlers.SetURL).Methods("GET")
	r.HandleFunc("/get/{key}", handlers.GetURL).Methods("GET")

	if strings.HasSuffix(*dataPath, ".yaml") {
		r.HandleFunc("/yaml/{key}", handlers.YamlHandler(*dataPath)).Methods("GET")
	}
	if strings.HasSuffix(*dataPath, ".json") {
		r.HandleFunc("/json/{key}", handlers.JsonHandler(*dataPath)).Methods("GET")
	}

	fmt.Println("Starting the server on :8080...")
	http.ListenAndServe(":8080", r)
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"urlShortener/handler"
)

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func main() {
	dataPath := flag.String("data", "data.yaml", "choose yaml file with data")
	flag.Parse()

	mux := defaultMux()

	Data, err := os.ReadFile(*dataPath)
	if err != nil {
		log.Fatalf("Unabale read file: %v\n", err)
	}


	if strings.HasSuffix(*dataPath, "yaml") {
		
		yamlHandler, err := urlshort.YAMLHandler([]byte(Data), mux)
		if err != nil {
			log.Fatalf("Unabale to parse yaml: %v\n", err)
		}
		fmt.Println("Starting the server on :8080...")
		http.ListenAndServe("127.0.0.1:8080", yamlHandler)

	} else  if strings.HasSuffix(*dataPath, "json") {

		jsonHandler, err := urlshort.JSONHandler([]byte(Data), mux)
		if err != nil {
			log.Fatalf("Unabale to parse yaml: %v\n", err)
		}
		fmt.Println("Starting the server on :8080...")
		http.ListenAndServe("127.0.0.1:8080", jsonHandler)
	}
}

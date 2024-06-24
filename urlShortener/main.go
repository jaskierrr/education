package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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
	yamlPath := flag.String("yaml", "data.yaml", "choose yaml file with data")
	flag.Parse()

	mux := defaultMux()


	yamlData, err := os.ReadFile(*yamlPath)
	if err != nil {
		log.Fatalf("Unabale read yaml file: %v\n", err)
	}

	yamlHandler, err := urlshort.YAMLHandler([]byte(yamlData), mux)
	if err != nil {
		log.Fatalf("Unabale to parse yaml: %v\n", err)
	}

	fmt.Println("Starting the server on :8080...")
	http.ListenAndServe("127.0.0.1:8080", yamlHandler)
}

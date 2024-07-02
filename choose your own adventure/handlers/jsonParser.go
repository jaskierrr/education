package handlers

import (
	"encoding/json"
	//"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
type Arc struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Story map[string]Arc



func parseJson() Story {
	file, err := os.Open("story.json")
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer file.Close()

	data := json.NewDecoder(file)
	if err != nil {
		log.Fatalf("Unable to parse JSON: %v", err)
	}

	var story Story

	if err = data.Decode(&story); err != nil {
		log.Fatalf("Unable to decode JSON: %v", err)
	}

	return story
}


func WriteStory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	story := parseJson()

	arc := story[key]


	tmpl, err := template.ParseFiles("main.html")
	if err != nil {
		log.Printf("Unable parse HTML tmpl: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, arc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Void(w http.ResponseWriter, r *http.Request)  {
	http.Redirect(w, r, "/intro", http.StatusPermanentRedirect)
}

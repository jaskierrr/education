package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"unicode"

	"github.com/gorilla/mux"
)

type Book struct {
	Intro     Intro     `json:"intro"`
	NewYork   NewYork   `json:"new-york"`
	Debate    Debate    `json:"debate"`
	SeanKelly SeanKelly `json:"sean-kelly"`
	MarkBates MarkBates `json:"mark-bates"`
	Denver    Denver    `json:"denver"`
	Home      Home      `json:"home"`
}
type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
type Intro struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}
type NewYork struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}
type Debate struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}
type SeanKelly struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}
type MarkBates struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}
type Denver struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}
type Home struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Options    `json:"options"`
}

func parseJson() Book {
	data, err := os.ReadFile("story.json")
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	book := Book{}

	err = json.Unmarshal(data, &book)
	if err != nil {
		log.Fatalf("Unable to parse JSON: %v", err)
	}

	return book
}

func titleGenerate(title string) reflect.Value {
	arr := strings.SplitAfter(title, "-")

	new := ""

	if len(arr) == 1 {
		runes := []rune(arr[0])

		runes[0] = unicode.ToUpper(runes[0])

		arr[0] = string(runes)

		new = arr[0]
	}

	if len(arr) == 2 {
		arr[0], _ = strings.CutSuffix(arr[0], "-")

		for i, v := range arr {
			runes := []rune(v)

			runes[0] = unicode.ToUpper(runes[0])

			arr[i] = string(runes)
		}

		new = arr[0] + arr[1]
	}

	fmt.Println(new)

	book := parseJson()

	field := reflect.ValueOf(book)

	value := field.FieldByName(new)

	return value

}

func WriteStory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	fmt.Println(key)

	title := titleGenerate(key)

	tmpl, err := template.ParseFiles("main.html")
	if err != nil {
		log.Printf("Unable parse HTML tmpl: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

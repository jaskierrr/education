package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
)

type yamlStruct struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

type jsonStruct struct {
	Path string `json:"path"`
	URL string `json:"url"`
}

func MapHandler(data map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]

		fmt.Println(key)

		if dest, ok := data[key]; ok {
			http.Redirect(w, r, dest, http.StatusPermanentRedirect)
		} else {
			fmt.Fprintf(w, "Unable redirect from: %q", key)
		}
	}
}

func YamlHandler(dataPath string) http.HandlerFunc  {
	data, err := os.ReadFile(dataPath)
	if err != nil {
		log.Fatalf("Unabale read file: %v\n", err)
	}

	parsedYaml, err := parseYaml(data)
	if err != nil {
		log.Fatalf("Unable parse YAML: %v", err)
	}

	return MapHandler(parsedYaml)
}

func parseYaml(data []byte) (map[string]string, error) {
	parsedYaml := []yamlStruct{}
	err := yaml.Unmarshal(data, &parsedYaml)

	yamlMap := make(map[string]string)
	for _, v := range parsedYaml {
		yamlMap[v.Path] = v.URL
	}

	return yamlMap, err
}

func JsonHandler(dataPath string) http.HandlerFunc  {
	data, err := os.ReadFile(dataPath)
	if err != nil {
		log.Fatalf("Unabale read file: %v\n", err)
	}

	parsedJson, err := parseJson(data)
	if err != nil {
		log.Fatalf("Unable parse JSON: %v", err)
	}

	return MapHandler(parsedJson)
}

func parseJson(data []byte) (map[string]string, error) {
	parsedYaml := []jsonStruct{}
	err := json.Unmarshal(data, &parsedYaml)

	yamlMap := make(map[string]string)
	for _, v := range parsedYaml {
		yamlMap[v.Path] = v.URL
	}
	return yamlMap, err
}

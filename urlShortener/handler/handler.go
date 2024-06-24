package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

type yamlStruct struct {
	Path string `yaml:"path"`
	URL string  `yaml:"url"`
}

type jsonStruct struct {
	Path string `json:"path"`
	URL string `json:"url"`
}

type yamlSlice []yamlStruct
type jsonSlice []jsonStruct

type MapBuilder interface {
	BuildMap()
}

func (slice yamlSlice) BuildMap() map[string]string {
	pathMap := make(map[string]string)
	for _, v := range slice {
		pathMap[v.Path] = v.URL
	}
	fmt.Println(pathMap)
	return pathMap
}

func (slice jsonSlice) BuildMap() map[string]string {
	pathMap := make(map[string]string)
	for _, v := range slice {
		pathMap[v.Path] = v.URL
	}
	fmt.Println(pathMap)
	return pathMap
}




// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if dest, ok := pathsToUrls[url]; ok {
			http.Redirect(w, r, dest, http.StatusPermanentRedirect)
		}
			fallback.ServeHTTP(w, r)
	})
}

// func buildMap(parsedData []yamlStruct) map[string]string {
// 	pathMap := make(map[string]string)
// 	for _, v := range parsedData {
// 		pathMap[v.Path] = v.URL
// 	}
// 	fmt.Println(pathMap)
// 	return pathMap
// }


// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYaml(data)
	if err != nil {
		return nil, err
	}

	pathMap := parsedYaml.BuildMap()

	return MapHandler(pathMap, fallback), nil
}

func parseYaml(data []byte) (yamlSlice, error) {
	parsedData := yamlSlice{}
	err := yaml.Unmarshal(data, &parsedData)

	return parsedData, err
}


func JSONHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJSON, err := parseJSON(data)
	if err != nil {
		return nil, err
	}

	pathMap := parsedJSON.BuildMap()

	return MapHandler(pathMap, fallback), nil
}

func parseJSON(data []byte) (jsonSlice, error) {
	parsedData := jsonSlice{}
	err := json.Unmarshal(data, &parsedData)

	return parsedData, err
}

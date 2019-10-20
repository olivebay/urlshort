package urlshort

import (
	"fmt"
	"net/http"

	"encoding/json"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
func YAMLHandler(yaml []byte) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yaml)

	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap), nil
}

// JSONHandler parses JSON data into a map, MapHandler then maps the path to the url
func JSONHandler(json []byte) (http.HandlerFunc, error) {
	parsedJSON, err := parseJSON(json)
	fmt.Println(string(json))
	if err != nil {
		return nil, err
	}
	pathMap := buildMapJSON(parsedJSON)
	return MapHandler(pathMap), nil
}

type mapper struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url"  json:"url"`
}

func parseYAML(data []byte) ([]mapper, error) {
	var parsedYAML []mapper

	err := yaml.Unmarshal(data, &parsedYAML)
	if err != nil {
		return nil, err
	}
	return parsedYAML, nil
}

func parseJSON(data []byte) ([]mapper, error) {
	var parsedJSON []mapper

	err := json.Unmarshal(data, &parsedJSON)
	if err != nil {
		return nil, err
	}
	return parsedJSON, nil
}

func buildMap(data []mapper) map[string]string {
	urlsMap := make(map[string]string)

	for _, s := range data {
		urlsMap[s.Path] = s.URL
	}
	return urlsMap
}

func buildMapJSON(data []mapper) map[string]string {
	urlsMap := make(map[string]string)

	for _, s := range data {
		urlsMap[s.Path] = s.URL
	}
	return urlsMap
}

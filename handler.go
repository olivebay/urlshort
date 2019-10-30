package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler map any paths (keys in the map) to their corresponding URL
// (values that each key in the map points to, in string format).
//
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
//
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
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
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yaml)

	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

// JSONHandler parses JSON data into a map, MapHandler then maps the path to the url
func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJSON, err := parseJSON(json)

	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedJSON)
	return MapHandler(pathMap, fallback), nil
}

type mapper struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url"  json:"url"`
}

// Unmarshal YAML into a mapper struct
func parseYAML(data []byte) ([]mapper, error) {
	var parsedYAML []mapper

	err := yaml.Unmarshal(data, &parsedYAML)
	if err != nil {
		return nil, err
	}
	return parsedYAML, nil
}

// Unmarshal JSON into a mapper struct
func parseJSON(data []byte) ([]mapper, error) {
	var parsedJSON []mapper

	err := json.Unmarshal(data, &parsedJSON)
	if err != nil {
		return nil, err
	}
	return parsedJSON, nil
}

// buildMap takes YAML or JSON data and builds a map of the data
func buildMap(data []mapper) map[string]string {
	urlsMap := make(map[string]string)

	for _, s := range data {
		urlsMap[s.Path] = s.URL
	}
	return urlsMap
}

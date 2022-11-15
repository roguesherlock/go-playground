package handler

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...

	return func(w http.ResponseWriter, r *http.Request) {
		path, ok := pathsToUrls[r.URL.Path]
		if ok {
			fmt.Println(path)
			fmt.Fprintln(w, path)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// parseYaml parses the yaml and returns a slice of pathUrl
func parseYaml(yml []byte) ([]pathUrl, error) {

	var data []pathUrl

	err := yaml.Unmarshal(yml, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// buildMap builds a map from the provided YAML data.
func buildMap(yml []pathUrl) map[string]string {
	data := make(map[string]string)
	for _, v := range yml {
		fmt.Println(v.Path, v.URL)
		data[v.Path] = v.URL
	}
	return data
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

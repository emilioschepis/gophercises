package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// The `yaml:"<something>"` part is referred to as "tags"
// and it allows us to define how to encode/decode the
// properties inside a struct
type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	// The signature of this function matches `http.HandlerFunc`
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// The `ok` parameter tells us if the map has the requested key
		// This essentially allows us to execute code only if the condition
		// is met
		// I have read that this pattern is called the "comma ok"
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

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
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}

	pathsToUrls := buildMap(pathUrls)

	return MapHandler(pathsToUrls, fallback), nil
}

func parseYaml(data []byte) ([]pathURL, error) {
	var pathUrls []pathURL

	// This "unmarshals" (aka decodes) the YAML bytes inside our slice
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}

	return pathUrls, nil
}

func buildMap(pathUrls []pathURL) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}
	return pathsToUrls
}

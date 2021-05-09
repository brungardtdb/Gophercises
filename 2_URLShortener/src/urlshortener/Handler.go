package urlshortener

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
	return func(res http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		if url, found := pathsToUrls[path]; found {
			http.Redirect(res, req, url, http.StatusSeeOther)
			return
		}
		fallback.ServeHTTP(res, req)
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

	yamlMap, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}

	umYamlMap, err := mapFromYaml(yamlMap)
	if err != nil {
		return nil, err
	}
	return MapHandler(umYamlMap, fallback), nil
}

type AppYaml struct {
	Path string
	URL  string
}

func parseYaml(yml []byte) ([]AppYaml, error) {

	var appYaml []AppYaml
	testYaml := yaml.Unmarshal(yml, &appYaml)

	if testYaml != nil {
		return nil, testYaml
	}
	return appYaml, nil
}

func mapFromYaml(umYaml []AppYaml) (map[string]string, error) {

	umYamlMap := make(map[string]string)
	fmt.Println(umYaml)

	for _, appYaml := range umYaml {
		umYamlMap[appYaml.Path] = appYaml.URL
	}
	return umYamlMap, nil
}

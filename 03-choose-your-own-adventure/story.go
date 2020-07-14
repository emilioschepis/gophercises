package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

// Story is the adventure
type Story map[string]Chapter

// Chapter is a portion of the adventure
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option is a choice that the user can make during the adventure
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type handler struct {
	s Story
}

var defaultHandlerTpl = `
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Choose Your Own Adventure</title>
</head>

<body>
  <!-- This refers to the "Title" property of cyoa.Story -->
  <h1>{{.Title}}</h1>
  <!-- This ranges through the paragraphs of the story -->
  {{range .Paragraphs}}
  <!-- Each paragraph is assigned to "." -->
  <p>{{.}}</p>
  {{end}}
  <ul>
    {{range .Options}}
    <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
    {{end}}
  </ul>
</body>

</html>
`

// NewHandler is...
// We return a generic `http.Handler` and not our `handler` struct
// to make this more generic and inherit all the docs of `http.Handler`
func NewHandler(s Story) http.Handler {
	return handler{s}
}

// JSONStory is...
func JSONStory(r io.Reader) (Story, error) {
	decoder := json.NewDecoder(r)
	var story Story
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

// This makes it so that our `handler` conforms to `http.Handler`
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace((r.URL.Path))
	if path == "" || path == "/" {
		path = "/intro"
	}

	// This removes the / prefix from the path
	// The operation is safe because if the path was empty it would
	// already have been replaced by "/intro"
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		tpl := template.Must(template.New("").Parse(defaultHandlerTpl))
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Chapter not found.", http.StatusNotFound)
}

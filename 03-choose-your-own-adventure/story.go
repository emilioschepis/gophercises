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
	t *template.Template
}

var tpl *template.Template

var defaultHandlerTpl = `
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Choose Your Own Adventure</title>
</head>

<body>
	<section class="page">
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
	</section>
	<style>
		body {
			font-family: helvetica, arial;
		}
		h1 {
			text-align:center;
			position:relative;
		}
		.page {
			width: 80%;
			max-width: 500px;
			margin: auto;
			margin-top: 40px;
			margin-bottom: 40px;
			padding: 80px;
			background: #FFFCF6;
			border: 1px solid #eee;
			box-shadow: 0 10px 6px -6px #777;
		}
		ul {
			border-top: 1px dotted #ccc;
			padding: 10px 0 0 0;
			-webkit-padding-start: 0;
		}
		li {
			padding-top: 10px;
		}
		a,
		a:visited {
			text-decoration: none;
			color: #6295b5;
		}
		a:active,
		a:hover {
			color: #7792a2;
		}
		p {
			text-indent: 1em;
		}
	</style>
</body>

</html>
`

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTpl))
}

// NewHandler is...
// We return a generic `http.Handler` and not our `handler` struct
// to make this more generic and inherit all the docs of `http.Handler`
func NewHandler(s Story, t *template.Template) http.Handler {
	if t == nil {
		t = tpl
	}

	return handler{s, t}
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
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Chapter not found.", http.StatusNotFound)
}

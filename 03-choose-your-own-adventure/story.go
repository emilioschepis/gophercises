package cyoa

import (
	"encoding/json"
	"io"
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

func JsonStory(r io.Reader) (Story, error) {
	decoder := json.NewDecoder(r)
	var story Story
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

package cyoa

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

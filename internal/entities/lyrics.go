package entities

type Lyrics struct {
	Id              int    `json:"id,omitempty"`
	Lyrics          string `json:"lyrics,omitempty"`
	FullLyrics      string `json:"fullLyrics,omitempty"`
	HasRights       bool   `json:"hasRights,omitempty"`
	ShowTranslation bool   `json:"showTranslation,omitempty"`
	TextLanguage    string `json:"textLanguage,omitempty"`
	Url             string `json:"url,omitempty"`
}

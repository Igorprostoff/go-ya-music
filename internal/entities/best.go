package entities

type Best struct {
	Type   string      `json:"type,omitempty"`
	Result interface{} `json:"result,omitempty"` //TODO: [Union[Track, Artist, Album, Playlist, Video]]
	Text   string      `json:"text,omitempty"`
}

package entities

type Event struct {
	Id          string        `json:"id,omitempty"`
	Type        string        `json:"type,omitempty"`
	TypeForFrom string        `json:"typeForFrom,omitempty"`
	Title       string        `json:"title,omitempty"`
	Tracks      []Track       `json:"tracks,omitempty"`
	Artists     []ArtistEvent `json:"artists,omitempty"`
	Albums      []AlbumEvent  `json:"albums,omitempty"`
	Message     string        `json:"message,omitempty"`
	Device      string        `json:"device,omitempty"`
	TracksCount int           `json:"tracksCount,omitempty"`
	Genre       string        `json:"genre,omitempty"`
}

package entities

type Counts struct {
	tracks       int `json:"tracks,omitempty"`
	directAlbums int `json:"directAlbums,omitempty"`
	alsoAlbums   int `json:"alsoAlbums,omitempty"`
	alsoTracks   int `json:"alsoTracks,omitempty"`
}

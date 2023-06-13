package entities

type ArtistEvent struct {
	Artist                      Artist   `json:"artist"`
	Tracks                      []Track  `json:"tracks,omitempty"`
	SimilarToArtistsFromHistory []Artist `json:"similarToArtistsFromHistory,omitempty"`
	Subscribed                  bool     `json:"subscribed,omitempty"`
}

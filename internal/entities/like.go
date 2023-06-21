package entities

type Like struct {
	Type             string   `json:"type,omitempty"`
	Id               string   `json:"id,omitempty"`
	Timestamp        string   `json:"timestamp,omitempty"`
	Album            Album    `json:"album"`
	Artist           Artist   `json:"artist"`
	Playlist         Playlist `json:"playlist"`
	ShortDescription string   `json:"shortDescription,omitempty"`
	Description      string   `json:"description,omitempty"`
	IsPremiere       bool     `json:"isPremiere,omitempty"`
	IsBanner         bool     `json:"isBanner,omitempty"`
}

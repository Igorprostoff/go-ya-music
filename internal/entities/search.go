package entities

type Search struct {
	SearchRequestId   string     `json:"searchRequestId,omitempty"`
	Text              string     `json:"text,omitempty"`
	Best              Best       `json:"best"`
	Albums            []Album    `json:"albums,omitempty"`
	Artists           []Artist   `json:"artists,omitempty"`
	Playlists         []Playlist `json:"playlists,omitempty"`
	Tracks            []Track    `json:"tracks,omitempty"`
	Videos            []Video    `json:"videos,omitempty"`
	Users             []User     `json:"users,omitempty"`
	Podcasts          []Album    `json:"podcasts,omitempty"`
	PodcastEpisodes   []Track    `json:"podcastEpisodes,omitempty"`
	Type              string     `json:"type,omitempty"`
	Page              int        `json:"page,omitempty"`
	PerPage           int        `json:"perPage,omitempty"`
	MisspellResult    string     `json:"misspellResult,omitempty"`
	MisspellOriginal  string     `json:"misspellOriginal,omitempty"`
	MisspellCorrected bool       `json:"misspellCorrected,omitempty"`
	Nocorrect         bool       `json:"nocorrect,omitempty"`
}

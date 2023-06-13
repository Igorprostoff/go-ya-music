package entities

type BriefInfo struct {
	Artist         Artist       `json:"artist"`
	Albums         []Album      `json:"albums,omitempty"`
	Playlists      []Playlist   `json:"playlists,omitempty"`
	AlsoAlbums     []Album      `json:"alsoAlbums,omitempty"`
	LastReleaseIds []int        `json:"lastReleaseIds,omitempty"`
	LastReleases   []Album      `json:"lastReleases,omitempty"`
	PopularTracks  []Track      `json:"popularTracks,omitempty"`
	SimilarArtists []Artist     `json:"similarArtists,omitempty"`
	AllCovers      []Cover      `json:"allCovers,omitempty"`
	Concerts       interface{}  `json:"concerts,omitempty"`
	Videos         []Video      `json:"videos,omitempty"`
	Vinyls         []Vinyl      `json:"vinyls,omitempty"`
	HasPromotions  bool         `json:"hasPromotions,omitempty"`
	PlaylistIds    []PlaylistId `json:"playlistIds,omitempty"`
	TracksInChart  []Chart      `json:"tracksInChart,omitempty"`
}

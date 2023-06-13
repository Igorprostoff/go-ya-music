package entities

type ArtistTracks struct {
	Tracks Track `json:"tracks"`
	Pager  Pager `json:"pager"`
}

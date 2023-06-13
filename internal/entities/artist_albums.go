package entities

type ArtistAlbums struct {
	Albums []Album `json:"albums,omitempty"`
	Pager  Pager   `json:"pager"`
}

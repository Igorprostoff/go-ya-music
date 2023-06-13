package entities

type AlbumEvent struct {
	Album  Album `json:"album"`
	Tracks Track `json:"tracks"`
}

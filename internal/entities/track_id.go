package entities

type TrackId struct {
	Id      int    `json:"id,omitempty"`
	TrackId int    `json:"trackId,omitempty"`
	AlbumId int    `json:"albumId,omitempty"`
	From_   string `json:"from_,omitempty"`
}

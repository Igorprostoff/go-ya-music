package entities

type Deprecation struct {
	TargetAlbumId int    `json:"targetAlbumId,omitempty"`
	Status        string `json:"status,omitempty"`
	Done          bool   `json:"done,omitempty"`
}

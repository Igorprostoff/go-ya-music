package entities

type Sequence struct {
	Type  string `json:"type,omitempty"`
	Track Track  `json:"track"`
	Liked bool   `json:"liked,omitempty"`
}

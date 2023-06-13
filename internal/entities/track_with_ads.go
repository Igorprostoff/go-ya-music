package entities

type TrackWithAds struct {
	Type  string `json:"type,omitempty"`
	Track Track  `json:"track"`
}

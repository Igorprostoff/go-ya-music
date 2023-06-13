package entities

type Day struct {
	Day                 string         `json:"day,omitempty"`
	Events              []Event        `json:"events,omitempty"`
	TracksToPlayWithAds []TrackWithAds `json:"tracksToPlayWithAds,omitempty"`
	TracksToPlay        []Track        `json:"tracksToPlay,omitempty"`
}

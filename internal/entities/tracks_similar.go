package entities

type TracksSimilar struct {
	Track         Track   `json:"track"`
	SimilarTracks []Track `json:"similarTracks,omitempty"`
}

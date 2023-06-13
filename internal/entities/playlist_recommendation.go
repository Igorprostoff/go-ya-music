package entities

type PlaylistRecommendation struct {
	Tracks  []Track `json:"tracks,omitempty"`
	BatchId string  `json:"batchId,omitempty"`
}

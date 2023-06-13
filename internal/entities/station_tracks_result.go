package entities

type StationTracksResult struct {
	Id       Id         `json:"id"`
	Sequence []Sequence `json:"sequence,omitempty"`
	BatchId  string     `json:"batchId,omitempty"`
	Pumpkin  bool       `json:"pumpkin,omitempty"`
}

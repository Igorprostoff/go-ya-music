package entities

type Queue struct {
	Context      Context   `json:"context"`
	Tracks       []TrackId `json:"tracks,omitempty"`
	CurrentIndex int       `json:"currentIndex,omitempty"`
	Modified     string    `json:"modified,omitempty"`
	Id           string    `json:"id,omitempty"`
	From_        string    `json:"from_,omitempty"`
}

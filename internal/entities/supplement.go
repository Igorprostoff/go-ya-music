package entities

type Supplement struct {
	Id               int             `json:"id,omitempty"`
	Lyrics           Lyrics          `json:"lyrics"`
	Videos           VideoSupplement `json:"videos"`
	RadioIsAvailable bool            `json:"radioIsAvailable,omitempty"`
	Description      string          `json:"description,omitempty"`
}

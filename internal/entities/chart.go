package entities

type Chart struct {
	Position  int     `json:"position,omitempty"`
	Progress  string  `json:"progress,omitempty"`
	Listeners int     `json:"listeners,omitempty"`
	Shift     int     `json:"shift,omitempty"`
	BgColor   string  `json:"bgColor,omitempty"`
	TrackId   TrackId `json:"trackId"`
}

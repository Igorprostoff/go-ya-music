package entities

type PlaylistAbsence struct {
	Kind   int    `json:"kind,omitempty"`
	Reason string `json:"reason,omitempty"`
}

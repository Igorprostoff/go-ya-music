package entities

type ShotData struct {
	CoverUri string   `json:"coverUri,omitempty"`
	MdsUrl   string   `json:"mdsUrl,omitempty"`
	ShotText string   `json:"shotText,omitempty"`
	ShotType ShotType `json:"shotType"`
}

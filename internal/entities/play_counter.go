package entities

type PlayCounter struct {
	Value       int    `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
	Updated     bool   `json:"updated,omitempty"`
}

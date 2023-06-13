package entities

type Suggestions struct {
	best        Best     `json:"best"`
	suggestions []string `json:"suggestions,omitempty"`
}

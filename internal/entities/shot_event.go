package entities

type shotEvent struct {
	EventId string `json:"eventId,omitempty"`
	Shots   []Shot `json:"shots,omitempty"`
}

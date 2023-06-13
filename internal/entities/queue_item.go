package entities

type QueueItem struct {
	Id       string  `json:"id,omitempty"`
	Context  Context `json:"context"`
	Modified string  `json:"modified,omitempty"`
}

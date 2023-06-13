package entities

type Contest struct {
	ContestId string `json:"contestId,omitempty"`
	Status    string `json:"status,omitempty"`
	CanEdit   bool   `json:"canEdit,omitempty"`
	Sent      string `json:"sent,omitempty"`
	Withdrawn string `json:"withdrawn,omitempty"`
}

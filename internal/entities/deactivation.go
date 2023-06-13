package entities

type Deactivation struct {
	Method       string `json:"method,omitempty"`
	Instructions string `json:"instructions,omitempty"`
}

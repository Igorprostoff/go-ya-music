package entities

type Enum struct {
	Type           string  `json:"type,omitempty"`
	Name           string  `json:"name,omitempty"`
	PossibleValues []Value `json:"possibleValues,omitempty"`
}

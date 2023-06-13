package entities

type DiscreteScale struct {
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	Min  Value  `json:"min"`
	Max  Value  `json:"max"`
}

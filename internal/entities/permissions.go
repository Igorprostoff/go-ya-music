package entities

type Permissions struct {
	Until   string   `json:"until,omitempty"`
	Values  []string `json:"values,omitempty"`
	Default []string `json:"default,omitempty"`
}

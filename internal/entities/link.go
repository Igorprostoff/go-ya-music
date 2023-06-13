package entities

type Link struct {
	Title         string `json:"title,omitempty"`
	Href          string `json:"href,omitempty"`
	Type          string `json:"type,omitempty"`
	CocialNetwork string `json:"cocialNetwork,omitempty"`
}

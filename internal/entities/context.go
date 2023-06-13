package entities

type Context struct {
	Type        string `json:"type,omitempty"`
	Id          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}

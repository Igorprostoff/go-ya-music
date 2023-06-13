package entities

type TagResult struct {
	Tag string       `json:"tag,omitempty"`
	Ids []PlaylistId `json:"ids,omitempty"`
}

package entities

type Landing struct {
	Pumpkin   bool    `json:"pumpkin,omitempty"`
	ContentId int     `json:"contentId,omitempty"` //string is possible
	Blocks    []Block `json:"blocks,omitempty"`
}

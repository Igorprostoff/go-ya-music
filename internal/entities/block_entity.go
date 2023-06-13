package entities

type BlockEntity struct {
	Id   string      `json:"id,omitempty"`
	Type string      `json:"type,omitempty"`
	Data interface{} `json:"data,omitempty"` //TODO: пока не понятно что нужно воткнуть
}

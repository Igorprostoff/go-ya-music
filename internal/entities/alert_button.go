package entities

type AlertButton struct {
	Text      string `json:"text,omitempty"`
	BgColor   string `json:"bgColor,omitempty"`
	TextColor string `json:"textColor,omitempty"`
	Uri       string `json:"uri,omitempty"`
}

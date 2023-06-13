package entities

type Alert struct {
	AlertId     string      `json:"alertId,omitempty"`
	Text        string      `json:"text,omitempty"`
	BgColor     string      `json:"bgColor,omitempty"`
	TextColor   string      `json:"textColor,omitempty"`
	AlertType   string      `json:"alertType,omitempty"`
	Button      AlertButton `json:"button"`
	CloseButton bool        `json:"closeButton,omitempty"`
}

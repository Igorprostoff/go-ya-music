package entities

type OpenGraphData struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Image       Cover  `json:"image"`
}

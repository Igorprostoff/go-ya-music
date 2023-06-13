package entities

type Brand struct {
	Image         string   `json:"image,omitempty"`
	Background    string   `json:"background,omitempty"`
	Reference     string   `json:"reference,omitempty"`
	Pixels        []string `json:"pixels,omitempty"`
	Theme         string   `json:"theme,omitempty"`
	PlaylistTheme []string `json:"playlistTheme,omitempty"`
	Button        []string `json:"button,omitempty"`
}

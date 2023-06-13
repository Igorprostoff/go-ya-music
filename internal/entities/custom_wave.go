package entities

type CustomWave struct {
	Title        string `json:"title,omitempty"`
	AnimationUrl string `json:"animationUrl,omitempty"`
	Position     string `json:"position,omitempty"`
}

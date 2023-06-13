package entities

type MixLink struct {
	Title              string `json:"title,omitempty"`
	Url                string `json:"url,omitempty"`
	UrlScheme          string `json:"urlScheme,omitempty"`
	TextColor          string `json:"textColor,omitempty"`
	BackgroundColor    string `json:"backgroundColor,omitempty"`
	BackgroundImageUri string `json:"backgroundImageUri,omitempty"`
	CoverWhite         string `json:"coverWhite,omitempty"`
	CoverUri           string `json:"coverUri,omitempty"`
}

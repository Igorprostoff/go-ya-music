package entities

type Promotion struct {
	PromoId   string `json:"promoId,omitempty"`
	Title     string `json:"title,omitempty"`
	Subtitle  string `json:"subtitle,omitempty"`
	Heading   string `json:"heading,omitempty"`
	Url       string `json:"url,omitempty"`
	UrlScheme string `json:"urlScheme,omitempty"`
	TextColor string `json:"textColor,omitempty"`
	Gradient  string `json:"gradient,omitempty"`
	Image     string `json:"image,omitempty"`
}

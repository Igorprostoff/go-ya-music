package entities

type VideoSupplement struct {
	Cover           string `json:"cover,omitempty"`
	Provider        string `json:"provider,omitempty"`
	Title           string `json:"title,omitempty"`
	ProviderVideoId string `json:"providerVideoId,omitempty"`
	Url             string `json:"url,omitempty"`
	EmbedUrl        string `json:"embedUrl,omitempty"`
	Embed           string `json:"embed,omitempty"`
}

package entities

type Video struct {
	title                   string   `json:"title,omitempty"`
	cover                   string   `json:"cover,omitempty"`
	embedUrl                string   `json:"embedUrl,omitempty"`
	provider                string   `json:"provider,omitempty"`
	providerVideoId         int      `json:"providerVideoId,omitempty"`
	youtubeUrl              string   `json:"youtubeUrl,omitempty"`
	thumbnailUrl            string   `json:"thumbnailUrl,omitempty"`
	duration                int      `json:"duration,omitempty"`
	text                    string   `json:"text,omitempty"`
	htmlAutoPlayVideoPlayer string   `json:"htmlAutoPlayVideoPlayer,omitempty"`
	regions                 []string `json:"regions,omitempty"`
}

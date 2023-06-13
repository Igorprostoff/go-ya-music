package entities

type GeneratedPlaylist struct {
	Type               string        `json:"type,omitempty"`
	Ready              bool          `json:"ready,omitempty"`
	Notify             bool          `json:"notify,omitempty"`
	Data               Playlist      `json:"data"`
	Description        []interface{} `json:"description,omitempty"`
	PreviewDescription string        `json:"previewDescription,omitempty"`
}

package entities

type Feed struct {
	CanGetMoreEvents   bool                `json:"canGetMoreEvents,omitempty"`
	Pumpkin            bool                `json:"pumpkin,omitempty"`
	IsWizardPassed     bool                `json:"isWizardPassed,omitempty"`
	GeneratedPlaylists []GeneratedPlaylist `json:"generatedPlaylists,omitempty"`
	Headlines          []interface{}       `json:"headlines,omitempty"`
	Today              string              `json:"today,omitempty"`
	Days               []Day               `json:"days,omitempty"`
	NextRevision       string              `json:"nextRevision,omitempty"`
}

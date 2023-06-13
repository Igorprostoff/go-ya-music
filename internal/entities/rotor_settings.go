package entities

type RotorSettings struct {
	Language   string `json:"language,omitempty"`
	Diversity  string `json:"diversity,omitempty"`
	Mood       int    `json:"mood,omitempty"`
	Energy     int    `json:"energy,omitempty"`
	MoodEnergy string `json:"moodEnergy,omitempty"`
}

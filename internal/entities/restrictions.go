package entities

type Restrictions struct {
	Language   Enum          `json:"language"`
	Diversity  Enum          `json:"diversity"`
	Mood       DiscreteScale `json:"mood"`
	Energy     DiscreteScale `json:"energy"`
	MoodEnergy Enum          `json:"moodEnergy"`
}

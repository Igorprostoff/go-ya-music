package entities

type CaseForms struct {
	Nominative    string `json:"nominative,omitempty"`
	Genitive      string `json:"genitive,omitempty"`
	Dative        string `json:"dative,omitempty"`
	Accusative    string `json:"accusative,omitempty"`
	Instrumental  string `json:"instrumental,omitempty"`
	Prepositional string `json:"prepositional,omitempty"`
}

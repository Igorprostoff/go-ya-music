package entities

type StationResult struct {
	station        Station       `json:"station"`
	settings       RotorSettings `json:"settings"`
	settings2      RotorSettings `json:"settings2"`
	adParams       AdParams      `json:"adParams"`
	explanation    string        `json:"explanation,omitempty"`
	prerolls       []interface{} `json:"prerolls,omitempty"`
	rupTitle       string        `json:"rupTitle,omitempty"`
	rupDescription string        `json:"rupDescription,omitempty"`
	customName     string        `json:"customName,omitempty"`
}

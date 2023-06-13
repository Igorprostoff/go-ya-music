package entities

type ChartInfo struct {
	Id               string        `json:"id,omitempty"`
	Type             string        `json:"type,omitempty"`
	TypeForFrom      string        `json:"typeForFrom,omitempty"`
	Title            string        `json:"title,omitempty"`
	Menu             ChartInfoMenu `json:"menu"`
	Chart            Playlist      `json:"chart"`
	ChartDescription string        `json:"chartDescription,omitempty"`
}

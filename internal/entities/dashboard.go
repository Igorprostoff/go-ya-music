package entities

type Dashboard struct {
	DashboardId string          `json:"dashboardId,omitempty"`
	Stations    []StationResult `json:"stations,omitempty"`
	Pumpkin     bool            `json:"pumpkin,omitempty"`
}

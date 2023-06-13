package entities

type Shot struct {
	order    int      `json:"order,omitempty"`
	played   bool     `json:"played,omitempty"`
	shotData ShotData `json:"shotData"`
	shotId   string   `json:"shotId,omitempty"`
	status   string   `json:"status,omitempty"`
}

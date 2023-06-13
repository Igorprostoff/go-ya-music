package entities

type NonAutoRenewable struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

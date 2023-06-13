package entities

type Subscription struct {
	NonAutoRenewableRemainder RenewableRemainder `json:"nonAutoRenewableRemainder"`
	AutoRenewable             []AutoRenewable    `json:"autoRenewable,omitempty"`
	FamilyAutoRenewable       []AutoRenewable    `json:"familyAutoRenewable,omitempty"`
	HadAnySubscription        bool               `json:"hadAnySubscription,omitempty"`
	Operator                  []Operator         `json:"operator,omitempty"`
	NonAutoRenewable          NonAutoRenewable   `json:"nonAutoRenewable"`
	CanStartTrial             bool               `json:"canStartTrial,omitempty"`
	Mcdonalds                 bool               `json:"mcdonalds,omitempty"`
	End                       string             `json:"end,omitempty"`
}

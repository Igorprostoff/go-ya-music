package entities

type Operator struct {
	ProductId         string         `json:"productId,omitempty"`
	Phone             string         `json:"phone,omitempty"`
	PaymentRegularity string         `json:"paymentRegularity,omitempty"`
	Deactivation      []Deactivation `json:"deactivation,omitempty"`
	Title             string         `json:"title,omitempty"`
	Suspended         bool           `json:"suspended,omitempty"`
}

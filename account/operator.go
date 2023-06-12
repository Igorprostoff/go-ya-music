package account

type Operator struct {
	Product_id         string
	Phone              string
	Payment_regularity string
	Deactivation       []Deactivation
	Title              string
	Suspended          bool
}

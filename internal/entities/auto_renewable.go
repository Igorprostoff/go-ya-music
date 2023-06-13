package entities

type AutoRenewable struct {
	Expires       string  `json:"expires,omitempty"`
	Vendor        string  `json:"vendor,omitempty"`
	VendorHelpUrl string  `json:"vendorHelpUrl,omitempty"`
	Product       Product `json:"product"`
	Finished      bool    `json:"finished,omitempty"`
	MasterInfo    User    `json:"masterInfo"`
	ProductId     string  `json:"productId,omitempty"`
	OrderId       int     `json:"orderId,omitempty"`
}

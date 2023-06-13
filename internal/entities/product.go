package entities

type Product struct {
	ProductId            string   `json:"productId,omitempty"`
	Type                 string   `json:"type,omitempty"`
	CommonPeriodDuration string   `json:"commonPeriodDuration,omitempty"`
	Duration             int      `json:"duration,omitempty"`
	TrialDuration        int      `json:"trialDuration,omitempty"`
	Price                Price    `json:"price"`
	Feature              string   `json:"feature,omitempty"`
	Debug                bool     `json:"debug,omitempty"`
	Plus                 bool     `json:"plus,omitempty"`
	Cheapest             bool     `json:"cheapest,omitempty"`
	Title                string   `json:"title,omitempty"`
	FamilySub            bool     `json:"familySub,omitempty"`
	FbImage              string   `json:"fbImage,omitempty"`
	FbName               string   `json:"fbName,omitempty"`
	Family               bool     `json:"family,omitempty"`
	Features             []string `json:"features,omitempty"`
	Description          string   `json:"description,omitempty"`
	Available            bool     `json:"available,omitempty"`
	TrialAvailable       bool     `json:"trialAvailable,omitempty"`
	TrialPeriodDuration  string   `json:"trialPeriodDuration,omitempty"`
	IntroPeriodDuration  string   `json:"introPeriodDuration,omitempty"`
	IntroPrice           Price    `json:"introPrice"`
	StartPeriodDuration  string   `json:"startPeriodDuration,omitempty"`
	StartPrice           Price    `json:"startPrice"`
	LicenceTextParts     Price    `json:"licenceTextParts"`
	VendorTrialAvailable bool     `json:"vendorTrialAvailable,omitempty"`
	ButtonText           string   `json:"buttonText,omitempty"`
	ButtonAdditionalText string   `json:"buttonAdditionalText,omitempty"`
	PaymentMethodTypes   []string `json:"paymentMethodTypes,omitempty"`
}

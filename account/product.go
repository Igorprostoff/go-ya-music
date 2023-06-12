package account

type Product struct {
	Product_id             string
	Type                   string
	Common_period_duration string
	Duration               int
	Trial_duration         int
	Price                  Price
	Feature                string
	Debug                  bool
	Plus                   bool
	Cheapest               bool
	Title                  string
	Family_sub             bool
	Fb_image               string
	Fb_name                string
	Family                 bool
	Features               []string
	Description            string
	Available              bool
	Trial_available        bool
	Trial_period_duration  string
	Intro_period_duration  string
	Intro_price            Price
	Start_period_duration  string
	Start_price            Price
	Licence_text_parts     Price
	Vendor_trial_available bool
	Button_text            string
	Button_additional_text string
	Payment_method_types   []string
}

package account

type Account struct {
	Now                      string
	Service_available        bool
	Region                   int
	Uid                      int
	Login                    string
	Full_name                string
	Second_name              string
	First_name               string
	Display_name             string
	Hosted_user              bool
	Birthday                 string
	Passport_phones          []Passport_phone
	Registered_at            string
	Has_info_for_app_metrica bool
	Child                    bool
}

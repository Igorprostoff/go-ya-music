package rotor

type Station_result struct {
	station         Station
	settings        Rotor_settings
	settings2       Rotor_settings
	ad_params       Ad_params
	explanation     string
	prerolls        []interface{}
	rup_title       string
	rup_description string
	custom_name     string
}

package rotor

type Station struct {
	Id                 Id
	Name               string
	Icon               Icon
	Mts_icon           Icon
	Geocell_icon       Icon
	Id_for_from        string
	Restrictions       Restrictions
	Restrictions2      Restrictions
	Full_image_url     string
	Mts_full_image_url string
	Parent_id          Id
}

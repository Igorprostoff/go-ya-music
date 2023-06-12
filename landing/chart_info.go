package landing

type Chart_info struct {
	Id                string
	Type              string
	Type_for_from     string
	Title             string
	Menu              Chart_info_menu
	Chart             Playlist
	Chart_description string
}

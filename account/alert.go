package account

type Alert struct {
	Alert_id     string
	Text         string
	Bg_color     string
	Text_color   string
	Alert_type   string
	Button       Alert_button
	Close_button bool
}

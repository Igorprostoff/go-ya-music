package genre

type Genre struct {
	Id              string
	Weight          int
	Composer_top    bool
	Title           string
	Titles          map[string]string
	Images          Images
	Show_in_menu    bool
	Show_in_regions []interface{}
	Full_title      string
	Url_part        string
	Color           string
	Radio_icon      Icon
	Sub_genres      []Genre
	Hide_in_regions interface{}
}

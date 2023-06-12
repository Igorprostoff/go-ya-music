package playlist

type User struct {
	Uid          int
	Login        string
	Name         string
	Display_name string
	Full_name    string
	Sex          string
	Verified     bool
	Regions      []int
}

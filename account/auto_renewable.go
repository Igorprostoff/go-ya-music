package account

import "github.com/Igorprostoff/go-ya-music/playlist"

type Auto_renewable struct {
	Expires         string
	Vendor          string
	Vendor_help_url string
	Product         Product
	Finished        bool
	Master_info     playlist.User
	Product_id      string
	Order_id        int
}

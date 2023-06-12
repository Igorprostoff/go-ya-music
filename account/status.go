package account

import "github.com/Igorprostoff/go-ya-music/rotor"

type Status struct {
	Account         Account
	Permissions     Permissions
	Advertisement   string
	Subscription    Subscription
	Cache_limit     int
	Subeditor       bool
	Subeditor_level int
	Plus            Plus
	Default_email   string
	Skips_per_hour  int
	Station_exists  bool
	Station_data    rotor.Station_data
	Bar_below       Alert
	Premium_region  int
	Experiment      int
	Pretrial_active bool
	Userhash        string
}

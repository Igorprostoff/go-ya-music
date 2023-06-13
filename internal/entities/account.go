package entities

type Account struct {
	Now                  string           `json:"now,omitempty"`
	ServiceAvailable     bool             `json:"serviceAvailable,omitempty"`
	Region               int              `json:"region,omitempty"`
	Uid                  int              `json:"uid,omitempty"`
	Login                string           `json:"login,omitempty"`
	FullName             string           `json:"fullName,omitempty"`
	SecondName           string           `json:"secondName,omitempty"`
	FirstName            string           `json:"firstName,omitempty"`
	DisplayName          string           `json:"displayName,omitempty"`
	HostedUser           bool             `json:"hostedUser,omitempty"`
	Birthday             string          `json:"birthday,omitempty"`
	PassportPhones       []PassportPhone `json:"passportPhones,omitempty"`
	RegisteredAt         string          `json:"registeredAt,omitempty"`
	HasInfoForAppMetrica bool             `json:"hasInfoForAppMetrica,omitempty"`
	Child                bool             `json:"child,omitempty"`
}

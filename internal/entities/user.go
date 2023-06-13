package entities

type User struct {
	Uid         int    `json:"uid,omitempty"`
	Login       string `json:"login,omitempty"`
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	FullName    string `json:"fullName,omitempty"`
	Sex         string `json:"sex,omitempty"`
	Verified    bool   `json:"verified,omitempty"`
	Regions     []int  `json:"regions,omitempty"`
}

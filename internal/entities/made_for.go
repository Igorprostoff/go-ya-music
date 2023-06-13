package entities

type MadeFor struct {
	UserInfo  User      `json:"userInfo"`
	CaseForms CaseForms `json:"caseForms"`
}

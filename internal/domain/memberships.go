package domain

type Membership struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Scopes      string `json:"-"`
}

type CompanyMembership struct {
	CompanyID    string `json:"company_id"`
	MembershipID string `json:"membership_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Scopes       string `json:"scopes"`
}

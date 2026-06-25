package model

// Domain is an asset network domain (v3 API) or zone (v4 API).
// AssetsAmount is populated in list responses; CreatedBy/UpdatedBy
// are returned by the API from the base model mixin.
type Domain struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Assets       IDNameList `json:"assets"`
	Gateways     IDNameList `json:"gateways"`
	AssetsAmount int        `json:"assets_amount,omitempty"`
	Comment      string     `json:"comment"`
	OrgID        string     `json:"org_id"`
	OrgName      string     `json:"org_name"`
	CreatedBy    string     `json:"created_by,omitempty"`
	UpdatedBy    string     `json:"updated_by,omitempty"`
	DateCreated   string     `json:"date_created"`
	DateUpdated   string     `json:"date_updated"`
}

// DomainRequest is the create/update payload.
type DomainRequest struct {
	ID       string   `json:"id,omitempty"`
	Name     string   `json:"name"`
	Assets   []string `json:"assets,omitempty"`
	Gateways []string `json:"gateways,omitempty"`
	Labels   []string `json:"labels,omitempty"`
	Comment  string   `json:"comment,omitempty"`
}

// DomainPage is the paginated list envelope.
type DomainPage struct {
	Total       int      `json:"count"`
	NextURL     string   `json:"next"`
	PreviousURL string   `json:"previous"`
	Results     []Domain `json:"results"`
}

package model

import "fmt"

// AssetCategory identifies an asset category.
type AssetCategory string

const (
	AssetCategoryHosts       AssetCategory = "hosts"
	AssetCategoryDevices     AssetCategory = "devices"
	AssetCategoryDatabases   AssetCategory = "databases"
	AssetCategoryWebs        AssetCategory = "webs"
	AssetCategoryClouds      AssetCategory = "clouds"
	AssetCategoryCustoms     AssetCategory = "customs"
	AssetCategoryDirectories AssetCategory = "directories"
)

func (c AssetCategory) String() string { return string(c) }

// Singular returns the category name without the trailing "s", used
// when building asset-type URLs.
func (c AssetCategory) Singular() string {
	s := string(c)
	if s == "" {
		return ""
	}
	if s == "databases" {
		return "database"
	}
	return s[:len(s)-1]
}

// AssetSpecInfo holds category-specific spec info.
type AssetSpecInfo struct {
	DBName           string           `json:"db_name,omitempty"`
	UseSSL           bool             `json:"use_ssl,omitempty"`
	AllowInvalidCert bool             `json:"allow_invalid_cert,omitempty"`
	Autofill         string           `json:"autofill,omitempty"`
	Script           []AssetWebScript `json:"script,omitempty"`
	SubmitSelector   string           `json:"submit_selector,omitempty"`
}

// AssetWebScript is a script entry for web assets.
type AssetWebScript struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Script any    `json:"script"`
}

// Asset is a generic JumpServer asset. The Domain field is populated on
// v3 responses; Zone is used on v4 responses. Both use any to handle
// the various shapes the API may return (object, string, or null).
type Asset struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Address      string        `json:"address"`
	Comment      string        `json:"comment"`
	Domain       any           `json:"domain,omitempty"`
	Zone         any           `json:"zone,omitempty"`
	Platform     PlatformMini  `json:"platform"`
	Nodes        IDNameList    `json:"nodes"`
	Labels       []any         `json:"labels"`
	Protocols    []any         `json:"protocols"`
	NodesDisplay []string      `json:"nodes_display"`
	Category     LabelValue    `json:"category"`
	Type         LabelValue    `json:"type"`
	Connectivity any           `json:"connectivity"`
	CreatedBy    string        `json:"created_by"`
	OrgID        string        `json:"org_id"`
	OrgName      string        `json:"org_name"`
	IsActive     bool          `json:"is_active"`
	DateVerified string        `json:"date_verified"`
	DateCreated  string        `json:"date_created"`
	SpecInfo     any           `json:"spec_info"`
}

// GetCategory returns the typed asset category.
func (a Asset) GetCategory() AssetCategory {
	if a.Category.Value == "ds" {
		return AssetCategoryDirectories
	}
	return AssetCategory(fmt.Sprintf("%ss", a.Category.Value))
}

// GetDomainOrZone returns a description of the domain/zone field,
// whichever is present. Returns nil if neither is set.
func (a Asset) GetDomainOrZone() any {
	if a.Zone != nil {
		return a.Zone
	}
	return a.Domain
}

// AssetRequest is the create/update payload. Set Zone for v4 or Domain
// for v3. When both are set, the service layer picks the right one
// based on the client version.
type AssetRequest struct {
	ID        string        `json:"id,omitempty"`
	Name      string        `json:"name"`
	Address   string        `json:"address"`
	Platform  int           `json:"platform"`
	Protocols []NamePort    `json:"protocols,omitempty"`
	Nodes     []string      `json:"nodes,omitempty"`
	Labels    []string      `json:"labels,omitempty"`
	Domain    string        `json:"domain,omitempty"`
	Zone      string        `json:"zone,omitempty"`
	IsActive  bool          `json:"is_active,omitempty"`
	Comment   string        `json:"comment,omitempty"`
	SpecInfo  AssetSpecInfo `json:"spec_info,omitempty"`
}

// AssetPage is the paginated list envelope.
type AssetPage struct {
	Total       int     `json:"count"`
	NextURL     string  `json:"next"`
	PreviousURL string  `json:"previous"`
	Results     []Asset `json:"results"`
}

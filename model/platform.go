package model

// PlatformProtocol is a rich protocol binding returned by the
// platform and gateway APIs, including settings and secret types.
type PlatformProtocol struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`
	Port         int            `json:"port"`
	PortFromAddr bool           `json:"port_from_addr"`
	Primary      bool           `json:"primary"`
	Required     bool           `json:"required"`
	Default      bool           `json:"default"`
	Public       bool           `json:"public"`
	SecretTypes  []string       `json:"secret_types"`
	Setting      map[string]any `json:"setting"`
}

// Platform describes a JumpServer platform template (Linux, Windows, etc.).
type Platform struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Category    LabelValue         `json:"category"`
	Type        LabelValue         `json:"type"`
	Charset     LabelValue         `json:"charset"`
	Internal    bool               `json:"internal"`
	Domain      bool               `json:"domain_enabled"`
	SuEnabled   bool               `json:"su_enabled"`
	Protocols   []PlatformProtocol `json:"protocols"`
	Comment     string             `json:"comment"`
	CreatedBy   string             `json:"created_by"`
	UpdatedBy   string             `json:"updated_by"`
	DateCreated string             `json:"date_created"`
	DateUpdated string             `json:"date_updated"`
}

// PlatformPage is the paginated list envelope for Platforms.
type PlatformPage = Page[Platform]

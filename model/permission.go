package model

// AssetPermission is an asset-authorization rule.
type AssetPermission struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Users       IDNameList   `json:"users"`
	UserGroups  IDNameList   `json:"user_groups"`
	Assets      IDNameList   `json:"assets"`
	Nodes       IDNameList   `json:"nodes"`
	Accounts    []string     `json:"accounts"`
	Protocols   []string     `json:"protocols"`
	Actions     []LabelValue `json:"actions"`
	IsActive    bool         `json:"is_active"`
	DateStart   string       `json:"date_start"`
	DateExpired string       `json:"date_expired"`
	Comment     string       `json:"comment"`
	OrgID       string       `json:"org_id"`
	OrgName     string       `json:"org_name"`
	CreatedBy   string       `json:"created_by"`
	DateCreated string       `json:"date_created"`
	DateUpdated string       `json:"date_updated"`
}

// AssetPermissionRequest is the create/update payload.
type AssetPermissionRequest struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name"`
	Users       []string `json:"users,omitempty"`
	UserGroups  []string `json:"user_groups,omitempty"`
	Assets      []string `json:"assets,omitempty"`
	Nodes       []string `json:"nodes,omitempty"`
	Accounts    []string `json:"accounts,omitempty"`
	Protocols   []string `json:"protocols,omitempty"`
	Actions     []string `json:"actions,omitempty"`
	IsActive    bool     `json:"is_active,omitempty"`
	DateStart   string   `json:"date_start,omitempty"`
	DateExpired string   `json:"date_expired,omitempty"`
	Comment     string   `json:"comment,omitempty"`
}

// AssetPermissionPage is the paginated list envelope for AssetPermissions.
type AssetPermissionPage = Page[AssetPermission]

// ---------- Asset permission relation (batch add) ----------

// AssetPermUserRelation is the request body for adding users to an asset permission.
type AssetPermUserRelation struct {
	User            string `json:"user"`
	AssetPermission string `json:"assetpermission"`
}

// AssetPermUserRelationDetail is the response for a user relation.
type AssetPermUserRelationDetail struct {
	ID                     int    `json:"id"`
	User                   string `json:"user"`
	UserDisplay            string `json:"user_display"`
	AssetPermission        string `json:"assetpermission"`
	AssetPermissionDisplay string `json:"assetpermission_display"`
}

// AssetPermUserGroupRelation is the request body for adding user groups to an asset permission.
type AssetPermUserGroupRelation struct {
	UserGroup       string `json:"usergroup"`
	AssetPermission string `json:"assetpermission"`
}

// AssetPermUserGroupRelationDetail is the response for a user group relation.
type AssetPermUserGroupRelationDetail struct {
	ID                     int    `json:"id"`
	UserGroup              string `json:"usergroup"`
	UserGroupDisplay       string `json:"usergroup_display"`
	AssetPermission        string `json:"assetpermission"`
	AssetPermissionDisplay string `json:"assetpermission_display"`
}

// AssetPermAssetRelation is the request body for adding assets to an asset permission.
type AssetPermAssetRelation struct {
	Asset           string `json:"asset"`
	AssetPermission string `json:"assetpermission"`
}

// AssetPermAssetRelationDetail is the response for an asset relation.
type AssetPermAssetRelationDetail struct {
	ID                     int    `json:"id"`
	Asset                  string `json:"asset"`
	AssetDisplay           string `json:"asset_display"`
	AssetPermission        string `json:"assetpermission"`
	AssetPermissionDisplay string `json:"assetpermission_display"`
}

// AssetPermNodeRelation is the request body for adding nodes to an asset permission.
type AssetPermNodeRelation struct {
	Node            string `json:"node"`
	AssetPermission string `json:"assetpermission"`
}

// AssetPermNodeRelationDetail is the response for a node relation.
type AssetPermNodeRelationDetail struct {
	ID                     int    `json:"id"`
	Node                   string `json:"node"`
	NodeDisplay            string `json:"node_display"`
	AssetPermission        string `json:"assetpermission"`
	AssetPermissionDisplay string `json:"assetpermission_display"`
}

// ---------- Self assets (my assets) ----------

// SelfAsset is an asset visible to the current user.
type SelfAsset struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Address      string       `json:"address"`
	Zone         any          `json:"zone"`
	Platform     PlatformMini `json:"platform"`
	OrgID        string       `json:"org_id"`
	Connectivity LabelValue   `json:"connectivity"`
	Nodes        []IDName     `json:"nodes"`
	Labels       []any        `json:"labels"`
	Category     LabelValue   `json:"category"`
	Type         LabelValue   `json:"type"`
	OrgName      string       `json:"org_name"`
	IsActive     bool         `json:"is_active"`
	DateVerified string       `json:"date_verified"`
	DateCreated  string       `json:"date_created"`
	Comment      string       `json:"comment"`
	CreatedBy    string       `json:"created_by"`
}

// SelfAssetPage is the paginated list envelope for SelfAssets.
type SelfAssetPage = Page[SelfAsset]

// PermedProtocolSetting contains protocol-specific settings.
type PermedProtocolSetting struct {
	Console     any    `json:"console"`
	Security    string `json:"security"`
	SftpHome    string `json:"sftp_home"`
	SftpEnabled bool   `json:"sftp_enabled"`
}

// PermedProtocol is a protocol available on a self asset.
type PermedProtocol struct {
	Name    string                `json:"name"`
	Port    int                   `json:"port"`
	Public  bool                  `json:"public"`
	Setting PermedProtocolSetting `json:"setting"`
}

// PermedAccount is an account available on a self asset.
type PermedAccount struct {
	ID          string       `json:"id"`
	Alias       string       `json:"alias"`
	Name        string       `json:"name"`
	Username    string       `json:"username"`
	HasUsername bool         `json:"has_username"`
	HasSecret   bool         `json:"has_secret"`
	SecretType  string       `json:"secret_type"`
	Actions     []LabelValue `json:"actions"`
	DateExpired string       `json:"date_expired"`
}

// SelfAssetDetail is the detailed view of a self asset, including
// permitted protocols and accounts.
type SelfAssetDetail struct {
	SelfAsset
	SpecInfo        any              `json:"spec_info"`
	PermedProtocols []PermedProtocol `json:"permed_protocols"`
	PermedAccounts  []PermedAccount  `json:"permed_accounts"`
}

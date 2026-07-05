package model

// PublicSetting is the subset of JumpServer settings exposed publicly.
type PublicSetting struct {
	Interface struct {
		LoginTitle  string `json:"login_title"`
		LogoURLs    any    `json:"logo_urls"`
		FaviconURLs any    `json:"favicon_urls"`
	} `json:"interface"`
	EnableWatermark          bool `json:"ENABLE_WATERMARK"`
	SecurityCommandExecution bool `json:"SECURITY_COMMAND_EXECUTION"`
	XpackLicenseIsValid      bool `json:"XPACK_LICENSE_IS_VALID"`
}

// SettingItem is a single setting entry.
type SettingItem struct {
	Name      string `json:"name"`
	Value     any    `json:"value"`
	Category  string `json:"category"`
	Encrypted bool   `json:"encrypted"`
}

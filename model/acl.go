package model

// CommandFilter is a command-filter ACL (v4).
type CommandFilter struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	CommandGroups any    `json:"command_groups"`
	Accounts      any    `json:"accounts"`
	Users         any    `json:"users"`
	UserGroups    any    `json:"user_groups"`
	Assets        any    `json:"assets"`
	Nodes         any    `json:"nodes"`
	Action        any    `json:"action"`
	IsActive      bool   `json:"is_active"`
	Priority      int    `json:"priority"`
	Comment       string `json:"comment"`
	OrgID         string `json:"org_id"`
	OrgName       string `json:"org_name"`
	DateCreated   string `json:"date_created"`
	DateUpdated   string `json:"date_updated"`
}

// CommandFilterRequest is the create/update payload. On v4, M2M fields
// (Users, Assets, etc.) accept either a list of IDs or a special object
// like {"type": "all"}.
type CommandFilterRequest struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name"`
	CommandGroups any    `json:"command_groups,omitempty"`
	Accounts      any    `json:"accounts,omitempty"`
	Users         any    `json:"users,omitempty"`
	UserGroups    any    `json:"user_groups,omitempty"`
	Assets        any    `json:"assets,omitempty"`
	Nodes         any    `json:"nodes,omitempty"`
	Action        string `json:"action"`
	IsActive      bool   `json:"is_active,omitempty"`
	Priority      int    `json:"priority,omitempty"`
	Comment       string `json:"comment,omitempty"`
}

// CommandFilterPage is the paginated list envelope.
type CommandFilterPage struct {
	Total       int             `json:"count"`
	NextURL     string          `json:"next"`
	PreviousURL string          `json:"previous"`
	Results     []CommandFilter `json:"results"`
}

// CommandGroup is a group of command regexes used by a CommandFilter.
type CommandGroup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        any    `json:"type"`
	Content     string `json:"content"`
	Comment     string `json:"comment"`
	OrgID       string `json:"org_id"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}

// CommandGroupRequest is the create/update payload.
type CommandGroupRequest struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name"`
	Type    any    `json:"type"`
	Content string `json:"content"`
	Comment string `json:"comment,omitempty"`
}

// CommandGroupPage is the paginated list envelope.
type CommandGroupPage struct {
	Total       int            `json:"count"`
	NextURL     string         `json:"next"`
	PreviousURL string         `json:"previous"`
	Results     []CommandGroup `json:"results"`
}

// LoginACL restricts login attempts by time/IP.
type LoginACL struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Action      LabelValue `json:"action"`
	IsActive    bool       `json:"is_active"`
	Priority    int        `json:"priority"`
	Comment     string     `json:"comment"`
	DateCreated string     `json:"date_created"`
	DateUpdated string     `json:"date_updated"`
}

// LoginACLPage is the paginated list envelope.
type LoginACLPage struct {
	Total       int        `json:"count"`
	NextURL     string     `json:"next"`
	PreviousURL string     `json:"previous"`
	Results     []LoginACL `json:"results"`
}

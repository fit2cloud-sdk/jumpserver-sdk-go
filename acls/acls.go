package acls

import (
	"context"

	"github.com/jumpserver-south/jumpserver-sdk-go/internal/core"
	"github.com/jumpserver-south/jumpserver-sdk-go/internal/sdkutil"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

// Command filter URL constants.
const (
	CommandFilterListURL   = "/api/v1/acls/command-filter-acls/"
	CommandFilterDetailURL = "/api/v1/acls/command-filter-acls/%s/"
	CommandGroupListURL    = "/api/v1/acls/command-groups/"
	CommandGroupDetailURL  = "/api/v1/acls/command-groups/%s/"
	CommandReviewURL       = "/api/v1/acls/command-filter-acls/command-review/"
)

// Login ACL URL constants.
const (
	LoginACLListURL    = "/api/v1/acls/login-acls/"
	LoginACLDetailURL  = "/api/v1/acls/login-acls/%s/"
	LoginAssetCheckURL = "/api/v1/acls/login-asset/check/"
)

// CommandFiltersService handles /api/v1/acls/command-filter-acls and
// /api/v1/acls/command-groups.
type CommandFiltersService struct {
	client core.HTTPClient
}

// NewCommandFiltersService creates a new CommandFiltersService.
func NewCommandFiltersService(c core.HTTPClient) *CommandFiltersService {
	return &CommandFiltersService{client: c}
}

// List returns a paginated list of command filters.
func (s *CommandFiltersService) List(ctx context.Context, opts *core.ListOptions) ([]model.CommandFilter, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(CommandFilterListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.CommandFilterPage
	resp, err := s.client.Do(ctx, httpReq, &page)
	if err != nil {
		return nil, resp, err
	}
	if resp != nil {
		resp.Count = page.Total
		resp.NextURL = page.NextURL
		resp.PreviousURL = page.PreviousURL
	}
	return page.Results, resp, nil
}

// Get fetches a command filter by ID.
func (s *CommandFiltersService) Get(ctx context.Context, id string) (*model.CommandFilter, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(CommandFilterDetailURL, id), nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.CommandFilter
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Create creates a command filter.
func (s *CommandFiltersService) Create(ctx context.Context, req *model.CommandFilterRequest) (*model.CommandFilter, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", CommandFilterListURL, req)
	if err != nil {
		return nil, nil, err
	}
	var out model.CommandFilter
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Update patches a command filter.
func (s *CommandFiltersService) Update(ctx context.Context, id string, req *model.CommandFilterRequest) (*model.CommandFilter, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "PATCH", sdkutil.Spath(CommandFilterDetailURL, id), req)
	if err != nil {
		return nil, nil, err
	}
	var out model.CommandFilter
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Delete deletes a command filter.
func (s *CommandFiltersService) Delete(ctx context.Context, id string) (*core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "DELETE", sdkutil.Spath(CommandFilterDetailURL, id), nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, httpReq, nil)
}

// ListGroups returns a paginated list of command groups.
func (s *CommandFiltersService) ListGroups(ctx context.Context, opts *core.ListOptions) ([]model.CommandGroup, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(CommandGroupListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.CommandGroupPage
	resp, err := s.client.Do(ctx, httpReq, &page)
	if err != nil {
		return nil, resp, err
	}
	if resp != nil {
		resp.Count = page.Total
		resp.NextURL = page.NextURL
		resp.PreviousURL = page.PreviousURL
	}
	return page.Results, resp, nil
}

// GetGroup fetches a command group by ID.
func (s *CommandFiltersService) GetGroup(ctx context.Context, id string) (*model.CommandGroup, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(CommandGroupDetailURL, id), nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.CommandGroup
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// CreateGroup creates a command group.
func (s *CommandFiltersService) CreateGroup(ctx context.Context, req *model.CommandGroupRequest) (*model.CommandGroup, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", CommandGroupListURL, req)
	if err != nil {
		return nil, nil, err
	}
	var out model.CommandGroup
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// UpdateGroup patches a command group.
func (s *CommandFiltersService) UpdateGroup(ctx context.Context, id string, req *model.CommandGroupRequest) (*model.CommandGroup, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "PATCH", sdkutil.Spath(CommandGroupDetailURL, id), req)
	if err != nil {
		return nil, nil, err
	}
	var out model.CommandGroup
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// DeleteGroup deletes a command group.
func (s *CommandFiltersService) DeleteGroup(ctx context.Context, id string) (*core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "DELETE", sdkutil.Spath(CommandGroupDetailURL, id), nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, httpReq, nil)
}

// LoginACLsService handles /api/v1/acls/login-acls.
type LoginACLsService struct {
	client core.HTTPClient
}

// NewLoginACLsService creates a new LoginACLsService.
func NewLoginACLsService(c core.HTTPClient) *LoginACLsService {
	return &LoginACLsService{client: c}
}

// List returns a paginated list of login ACLs.
func (s *LoginACLsService) List(ctx context.Context, opts *core.ListOptions) ([]model.LoginACL, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(LoginACLListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.LoginACLPage
	resp, err := s.client.Do(ctx, httpReq, &page)
	if err != nil {
		return nil, resp, err
	}
	if resp != nil {
		resp.Count = page.Total
		resp.NextURL = page.NextURL
		resp.PreviousURL = page.PreviousURL
	}
	return page.Results, resp, nil
}

// Get fetches a login ACL by ID.
func (s *LoginACLsService) Get(ctx context.Context, id string) (*model.LoginACL, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(LoginACLDetailURL, id), nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.LoginACL
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

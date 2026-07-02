package rbac

import (
	"context"

	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/core"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/sdkutil"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/model"
)

const (
	RoleListURL   = "/api/v1/rbac/%s-roles/"
	RoleDetailURL = "/api/v1/rbac/%s-roles/%s/"
)

// Service handles /api/v1/rbac/{scope}-roles.
type Service struct {
	client core.HTTPClient
}

// NewService creates a new RBAC Service.
func NewService(c core.HTTPClient) *Service {
	return &Service{client: c}
}

// List returns a paginated list of roles. Scope is "org" or "system".
func (s *Service) List(ctx context.Context, scope string, opts *core.ListOptions) ([]model.Role, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(sdkutil.Spath(RoleListURL, scope), params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.RolePage
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

// Get fetches a role by scope + ID.
func (s *Service) Get(ctx context.Context, scope, id string) (*model.Role, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(RoleDetailURL, scope, id), nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.Role
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

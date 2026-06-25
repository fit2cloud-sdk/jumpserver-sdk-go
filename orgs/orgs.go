package orgs

import (
	"context"

	"github.com/jumpserver-south/jumpserver-sdk-go/internal/core"
	"github.com/jumpserver-south/jumpserver-sdk-go/internal/sdkutil"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

const (
	ListURL   = "/api/v1/orgs/orgs/"
	DetailURL = "/api/v1/orgs/orgs/%s/"
)

// Service handles /api/v1/orgs/orgs.
type Service struct {
	client core.HTTPClient
}

// NewService creates a new orgs Service.
func NewService(c core.HTTPClient) *Service {
	return &Service{client: c}
}

// List returns a paginated list of organizations.
func (s *Service) List(ctx context.Context, opts *core.ListOptions) ([]model.Organization, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(ListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.OrganizationPage
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

// Get fetches an organization by ID.
func (s *Service) Get(ctx context.Context, id string) (*model.Organization, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(DetailURL, id), nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.Organization
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Create creates an organization.
func (s *Service) Create(ctx context.Context, req *model.OrganizationRequest) (*model.Organization, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", ListURL, req)
	if err != nil {
		return nil, nil, err
	}
	var out model.Organization
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Update patches an organization.
func (s *Service) Update(ctx context.Context, id string, req *model.OrganizationRequest) (*model.Organization, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "PATCH", sdkutil.Spath(DetailURL, id), req)
	if err != nil {
		return nil, nil, err
	}
	var out model.Organization
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Delete deletes an organization.
func (s *Service) Delete(ctx context.Context, id string) (*core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "DELETE", sdkutil.Spath(DetailURL, id), nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, httpReq, nil)
}

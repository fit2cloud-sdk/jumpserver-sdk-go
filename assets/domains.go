package assets

import (
	"context"

	"github.com/jumpserver-south/jumpserver-sdk-go/internal/core"
	"github.com/jumpserver-south/jumpserver-sdk-go/internal/sdkutil"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

// Domains URL constants.
const (
	DomainListURL   = "/api/v1/assets/domains/"
	DomainDetailURL = "/api/v1/assets/domains/%s/"
)

// Deprecated: Use ZonesService instead.
// DomainsService handles /api/v1/assets/domains (v3).
type DomainsService struct {
	client core.HTTPClient
}

// Deprecated: Use NewZonesService instead.
// NewDomainsService creates a new DomainsService.
func NewDomainsService(c core.HTTPClient) *DomainsService {
	return &DomainsService{client: c}
}

// List returns a paginated list of domains.
func (s *DomainsService) List(ctx context.Context, opts *core.ListOptions) ([]model.Domain, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(DomainListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.DomainPage
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

// Get fetches a domain by ID.
func (s *DomainsService) Get(ctx context.Context, id string) (*model.Domain, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(DomainDetailURL, id), nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.Domain
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Create creates a domain.
func (s *DomainsService) Create(ctx context.Context, req *model.DomainRequest) (*model.Domain, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", DomainListURL, req)
	if err != nil {
		return nil, nil, err
	}
	var out model.Domain
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Update patches a domain.
func (s *DomainsService) Update(ctx context.Context, id string, req *model.DomainRequest) (*model.Domain, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "PATCH", sdkutil.Spath(DomainDetailURL, id), req)
	if err != nil {
		return nil, nil, err
	}
	var out model.Domain
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Delete deletes a domain.
func (s *DomainsService) Delete(ctx context.Context, id string) (*core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "DELETE", sdkutil.Spath(DomainDetailURL, id), nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, httpReq, nil)
}

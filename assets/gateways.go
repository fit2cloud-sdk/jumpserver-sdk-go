package assets

import (
	"context"

	"github.com/jumpserver-south/jumpserver-sdk-go/internal/core"
	"github.com/jumpserver-south/jumpserver-sdk-go/internal/sdkutil"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

// Gateways URL constants.
const (
	GatewayListURL   = "/api/v1/assets/gateways/"
	GatewayDetailURL = "/api/v1/assets/gateways/%s/"
)

// GatewaysService handles /api/v1/assets/gateways.
type GatewaysService struct {
	client core.HTTPClient
}

// NewGatewaysService creates a new GatewaysService.
func NewGatewaysService(c core.HTTPClient) *GatewaysService {
	return &GatewaysService{client: c}
}

// List returns a paginated list of gateways.
func (s *GatewaysService) List(ctx context.Context, opts *core.ListOptions) ([]model.Gateway, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(GatewayListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.GatewayPage
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

// Get fetches a gateway by ID.
func (s *GatewaysService) Get(ctx context.Context, id string) (*model.Gateway, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(GatewayDetailURL, id), nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.Gateway
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Create creates a gateway.
func (s *GatewaysService) Create(ctx context.Context, req *model.GatewayRequest) (*model.Gateway, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", GatewayListURL, req)
	if err != nil {
		return nil, nil, err
	}
	var out model.Gateway
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Update patches a gateway.
func (s *GatewaysService) Update(ctx context.Context, id string, req *model.GatewayRequest) (*model.Gateway, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "PATCH", sdkutil.Spath(GatewayDetailURL, id), req)
	if err != nil {
		return nil, nil, err
	}
	var out model.Gateway
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Delete deletes a gateway.
func (s *GatewaysService) Delete(ctx context.Context, id string) (*core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "DELETE", sdkutil.Spath(GatewayDetailURL, id), nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, httpReq, nil)
}

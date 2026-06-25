package assets

import (
	"context"

	"github.com/jumpserver-south/jumpserver-sdk-go/internal/core"
	"github.com/jumpserver-south/jumpserver-sdk-go/internal/sdkutil"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

// Zones URL constants.
const (
	ZoneListURL   = "/api/v1/assets/zones/"
	ZoneDetailURL = "/api/v1/assets/zones/%s/"
)

// ZonesService handles network zones on v4 and automatically falls
// back to /api/v1/assets/domains on v3. Callers should always use
// ZonesService; the correct endpoint is chosen based on the client's
// configured version.
type ZonesService struct {
	client core.HTTPClient
}

// NewZonesService creates a new ZonesService.
func NewZonesService(c core.HTTPClient) *ZonesService {
	return &ZonesService{client: c}
}

func (s *ZonesService) listURL() string {
	if s.client.Version() == core.JumpServerV3.String() {
		return DomainListURL
	}
	return ZoneListURL
}

func (s *ZonesService) detailURL() string {
	if s.client.Version() == core.JumpServerV3.String() {	
		return DomainDetailURL
	}
	return ZoneDetailURL
}

// List returns a paginated list of zones (v4) or domains (v3).
func (s *ZonesService) List(ctx context.Context, opts *core.ListOptions) ([]model.Zone, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(s.listURL(), params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.ZonePage
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

// Get fetches a zone (v4) or domain (v3) by ID.
func (s *ZonesService) Get(ctx context.Context, id string) (*model.Zone, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(s.detailURL(), id), nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.Zone
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Create creates a zone (v4) or domain (v3).
func (s *ZonesService) Create(ctx context.Context, req *model.ZoneRequest) (*model.Zone, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", s.listURL(), req)
	if err != nil {
		return nil, nil, err
	}
	var out model.Zone
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Update patches a zone (v4) or domain (v3).
func (s *ZonesService) Update(ctx context.Context, id string, req *model.ZoneRequest) (*model.Zone, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "PATCH", sdkutil.Spath(s.detailURL(), id), req)
	if err != nil {
		return nil, nil, err
	}
	var out model.Zone
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Delete deletes a zone (v4) or domain (v3) by ID.
func (s *ZonesService) Delete(ctx context.Context, id string) (*core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "DELETE", sdkutil.Spath(s.detailURL(), id), nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, httpReq, nil)
}

package assets

import (
	"context"

	"github.com/jumpserver-south/jumpserver-sdk-go/internal/core"
	"github.com/jumpserver-south/jumpserver-sdk-go/internal/sdkutil"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

// Platforms URL constants.
const (
	PlatformListURL   = "/api/v1/assets/platforms/"
	PlatformDetailURL = "/api/v1/assets/platforms/%d/"
	ProtocolListURL   = "/api/v1/assets/protocols/"
)

// PlatformsService handles /api/v1/assets/platforms.
type PlatformsService struct {
	client core.HTTPClient
}

// NewPlatformsService creates a new PlatformsService.
func NewPlatformsService(c core.HTTPClient) *PlatformsService {
	return &PlatformsService{client: c}
}

// List returns a paginated list of platforms.
func (s *PlatformsService) List(ctx context.Context, opts *core.ListOptions) ([]model.Platform, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(PlatformListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.PlatformPage
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

// Get fetches a platform by ID.
func (s *PlatformsService) Get(ctx context.Context, id int) (*model.Platform, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(PlatformDetailURL, id), nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.Platform
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

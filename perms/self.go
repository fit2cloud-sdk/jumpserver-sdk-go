package perms

import (
	"context"

	"github.com/jumpserver-south/jumpserver-sdk-go/internal/core"
	"github.com/jumpserver-south/jumpserver-sdk-go/internal/sdkutil"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

// SelfService handles endpoints for the current user's own resources
// (assets, accounts, etc.).
type SelfService struct {
	client core.HTTPClient
}

// NewSelfService creates a new SelfService.
func NewSelfService(c core.HTTPClient) *SelfService {
	return &SelfService{client: c}
}

// ListAssets returns a paginated list of assets visible to the current user.
func (s *SelfService) ListAssets(ctx context.Context, filters map[string]string, opts *core.ListOptions) ([]model.SelfAsset, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	for k, v := range filters {
		if v != "" {
			params[k] = v
		}
	}
	path := sdkutil.AppendQuery(SelfAssetsListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.SelfAssetPage
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

// GetAsset returns the detailed view of an asset visible to the current user.
func (s *SelfService) GetAsset(ctx context.Context, assetID string) (*model.SelfAssetDetail, *core.Response, error) {
	return sdkutil.Get[model.SelfAssetDetail](ctx, s.client, SelfAssetAccountsURL, assetID)
}

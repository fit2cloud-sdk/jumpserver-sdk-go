package assets

import (
	"context"

	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/core"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/util"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/model"
)

// Assets URL constants.
const (
	AssetListURL          = "/api/v1/assets/assets/"
	AssetDetailURL        = "/api/v1/assets/assets/%s/"
	AssetPermUsersURL     = "/api/v1/assets/assets/%s/perm-users/"
	AssetPermUserPermsURL = "/api/v1/assets/assets/%s/perm-users/%s/permissions/"
)

// AssetsService handles the generic /api/v1/assets/assets endpoints.
type AssetsService struct {
	client core.HTTPClient
}

// NewAssetsService creates a new AssetsService.
func NewAssetsService(c core.HTTPClient) *AssetsService {
	return &AssetsService{client: c}
}

// List returns a paginated list of assets. Pass nil filters for no
// resource-specific filtering; common pagination goes in opts.
func (s *AssetsService) List(ctx context.Context, filters map[string]string, opts *core.ListOptions) ([]model.Asset, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	for k, v := range filters {
		if v != "" {
			params[k] = v
		}
	}
	return util.ListWithParams[model.Asset](ctx, s.client, AssetListURL, params)
}

// Get fetches an asset by ID.
func (s *AssetsService) Get(ctx context.Context, id string) (*model.Asset, *core.Response, error) {
	return util.Get[model.Asset](ctx, s.client, AssetDetailURL, id)
}

// Delete deletes an asset by ID.
func (s *AssetsService) Delete(ctx context.Context, id string) (*core.Response, error) {
	return util.Delete(ctx, s.client, AssetDetailURL, id)
}
// BatchDelete deletes multiple assets by ID using the cache-then-delete pattern.
func (s *AssetsService) BatchDelete(ctx context.Context, ids []string) error {
	return util.BatchDelete(ctx, s.client, AssetListURL, ids)
}

// PermUsers returns the users permitted to access an asset.
func (s *AssetsService) PermUsers(ctx context.Context, assetID string, opts *core.ListOptions) ([]model.User, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := util.AppendQuery(util.Spath(AssetPermUsersURL, assetID), params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var out []model.User
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return out, resp, nil
}

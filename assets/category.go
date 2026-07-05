package assets

import (
	"context"

	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/core"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/util"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/model"
)

// CategoryService is a typed facade over a single asset category
// (hosts, devices, databases, webs, clouds, customs).
type CategoryService struct {
	client    core.HTTPClient
	listURL   string
	detailURL string
}

// NewCategoryService creates a new CategoryService for the given category.
func NewCategoryService(c core.HTTPClient, category string) *CategoryService {
	return &CategoryService{
		client:    c,
		listURL:   "/api/v1/assets/" + category + "s/",
		detailURL: "/api/v1/assets/" + category + "s/%s/",
	}
}

// List returns a paginated list of assets in this category.
func (s *CategoryService) List(ctx context.Context, filters map[string]string, opts *core.ListOptions) ([]model.Asset, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	for k, v := range filters {
		if v != "" {
			params[k] = v
		}
	}
	return util.ListWithParams[model.Asset](ctx, s.client, s.listURL, params)
}

// Get fetches a category-scoped asset.
func (s *CategoryService) Get(ctx context.Context, id string) (*model.Asset, *core.Response, error) {
	return util.Get[model.Asset](ctx, s.client, s.detailURL, id)
}

// Create creates a category-scoped asset.
func (s *CategoryService) Create(ctx context.Context, req *model.AssetRequest) (*model.Asset, *core.Response, error) {
	return util.Create[model.Asset, model.AssetRequest](ctx, s.client, s.listURL, req)
}

// Update patches a category-scoped asset.
func (s *CategoryService) Update(ctx context.Context, id string, req *model.AssetRequest) (*model.Asset, *core.Response, error) {
	return util.Update[model.Asset, model.AssetRequest](ctx, s.client, s.detailURL, id, req)
}

// Replace replaces a category-scoped asset.
func (s *CategoryService) Replace(ctx context.Context, id string, req *model.AssetRequest) (*model.Asset, *core.Response, error) {
	return util.Replace[model.Asset, model.AssetRequest](ctx, s.client, s.detailURL, id, req)
}

// Delete deletes a category-scoped asset.
func (s *CategoryService) Delete(ctx context.Context, id string) (*core.Response, error) {
	return util.Delete(ctx, s.client, s.detailURL, id)
}

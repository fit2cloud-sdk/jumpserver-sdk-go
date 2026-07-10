package assets

import (
	"context"

	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/core"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/util"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/model"
)

// Zones URL constants.
const (
	ZoneListURL   = "/api/v1/assets/zones/"
	ZoneDetailURL = "/api/v1/assets/zones/%s/"
)

// ZonesService handles /api/v1/assets/zones network zones.
type ZonesService struct {
	client core.HTTPClient
}

// NewZonesService creates a new ZonesService.
func NewZonesService(c core.HTTPClient) *ZonesService {
	return &ZonesService{client: c}
}

// List returns a paginated list of zones.
func (s *ZonesService) List(ctx context.Context, opts *core.ListOptions) ([]model.Zone, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	return util.List[model.Zone](ctx, s.client, ZoneListURL, opts)
}

// Get fetches a zone by ID.
func (s *ZonesService) Get(ctx context.Context, id string) (*model.Zone, *core.Response, error) {
	return util.Get[model.Zone](ctx, s.client, ZoneDetailURL, id)
}

// Create creates a zone.
func (s *ZonesService) Create(ctx context.Context, req *model.ZoneRequest) (*model.Zone, *core.Response, error) {
	return util.Create[model.Zone, model.ZoneRequest](ctx, s.client, ZoneListURL, req)
}

// Update patches a zone.
func (s *ZonesService) Update(ctx context.Context, id string, req *model.ZoneRequest) (*model.Zone, *core.Response, error) {
	return util.Update[model.Zone, model.ZoneRequest](ctx, s.client, ZoneDetailURL, id, req)
}

// Delete deletes a zone by ID.
func (s *ZonesService) Delete(ctx context.Context, id string) (*core.Response, error) {
	return util.Delete(ctx, s.client, ZoneDetailURL, id)
}
// BatchDelete deletes multiple zones by ID using the cache-then-delete pattern.
func (s *ZonesService) BatchDelete(ctx context.Context, ids []string) error {
	return util.BatchDelete(ctx, s.client, ZoneListURL, ids)
}

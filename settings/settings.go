package settings

import (
	"context"

	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/core"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/util"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/model"
)

const (
	PublicSettingURL = "/api/v1/settings/public/"
	SettingListURL   = "/api/v1/settings/setting/"
)

// Service handles /api/v1/settings.
type Service struct {
	client core.HTTPClient
}

// NewService creates a new settings Service.
func NewService(c core.HTTPClient) *Service {
	return &Service{client: c}
}

// Public returns the public settings blob.
func (s *Service) Public(ctx context.Context) (*model.PublicSetting, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", PublicSettingURL, nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.PublicSetting
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// List returns the full settings map (requires admin). The v4 API
// returns a JSON object keyed by setting name, not an array.
func (s *Service) List(ctx context.Context, opts *core.ListOptions) (map[string]any, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := util.AppendQuery(SettingListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	out := map[string]any{}
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return out, resp, nil
}

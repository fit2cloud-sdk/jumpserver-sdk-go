package accounts

import (
	"context"

	"github.com/jumpserver-south/jumpserver-sdk-go/internal/core"
	"github.com/jumpserver-south/jumpserver-sdk-go/internal/sdkutil"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

// Account template URL constants.
const (
	TemplateListURL   = "/api/v1/accounts/account-templates/"
	TemplateDetailURL = "/api/v1/accounts/account-templates/%s/"
)

// TemplatesService handles /api/v1/accounts/account-templates.
type TemplatesService struct {
	client core.HTTPClient
}

// NewTemplatesService creates a new TemplatesService.
func NewTemplatesService(c core.HTTPClient) *TemplatesService {
	return &TemplatesService{client: c}
}

// List returns a paginated list of account templates.
func (s *TemplatesService) List(ctx context.Context, opts *core.ListOptions) ([]model.AccountTemplate, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(TemplateListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.AccountTemplatePage
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

// Get fetches an account template by ID.
func (s *TemplatesService) Get(ctx context.Context, id string) (*model.AccountTemplate, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(TemplateDetailURL, id), nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.AccountTemplate
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Create creates an account template.
func (s *TemplatesService) Create(ctx context.Context, req *model.AccountTemplateRequest) (*model.AccountTemplate, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", TemplateListURL, req)
	if err != nil {
		return nil, nil, err
	}
	var out model.AccountTemplate
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Update patches an account template.
func (s *TemplatesService) Update(ctx context.Context, id string, req *model.AccountTemplateRequest) (*model.AccountTemplate, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "PATCH", sdkutil.Spath(TemplateDetailURL, id), req)
	if err != nil {
		return nil, nil, err
	}
	var out model.AccountTemplate
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Delete deletes an account template.
func (s *TemplatesService) Delete(ctx context.Context, id string) (*core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "DELETE", sdkutil.Spath(TemplateDetailURL, id), nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, httpReq, nil)
}

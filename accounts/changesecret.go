package accounts

import (
	"context"

	"github.com/jumpserver-south/jumpserver-sdk-go/internal/core"
	"github.com/jumpserver-south/jumpserver-sdk-go/internal/sdkutil"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

// Change secret automation URL constants.
const (
	ChangeSecretListURL    = "/api/v1/accounts/change-secret-automations/"
	ChangeSecretDetailURL  = "/api/v1/accounts/change-secret-automations/%s/"
	ChangeSecretExecuteURL = "/api/v1/accounts/change-secret-executions/"
)

// ChangeSecretService handles /api/v1/accounts/change-secret-automations.
type ChangeSecretService struct {
	client core.HTTPClient
}

// NewChangeSecretService creates a new ChangeSecretService.
func NewChangeSecretService(c core.HTTPClient) *ChangeSecretService {
	return &ChangeSecretService{client: c}
}

// List returns a paginated list of change secret automations.
func (s *ChangeSecretService) List(ctx context.Context, opts *core.ListOptions) ([]model.ChangeSecretAutomation, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(ChangeSecretListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.ChangeSecretAutomationPage
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

// Get fetches a change secret automation by ID.
func (s *ChangeSecretService) Get(ctx context.Context, id string) (*model.ChangeSecretAutomation, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(ChangeSecretDetailURL, id), nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.ChangeSecretAutomation
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Create creates a change secret automation.
func (s *ChangeSecretService) Create(ctx context.Context, req *model.ChangeSecretAutomationRequest) (*model.ChangeSecretAutomation, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", ChangeSecretListURL, req)
	if err != nil {
		return nil, nil, err
	}
	var out model.ChangeSecretAutomation
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Update patches a change secret automation.
func (s *ChangeSecretService) Update(ctx context.Context, id string, req *model.ChangeSecretAutomationRequest) (*model.ChangeSecretAutomation, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "PATCH", sdkutil.Spath(ChangeSecretDetailURL, id), req)
	if err != nil {
		return nil, nil, err
	}
	var out model.ChangeSecretAutomation
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Delete deletes a change secret automation.
func (s *ChangeSecretService) Delete(ctx context.Context, id string) (*core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "DELETE", sdkutil.Spath(ChangeSecretDetailURL, id), nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, httpReq, nil)
}

// Execute triggers a change secret execution for the given automation.
func (s *ChangeSecretService) Execute(ctx context.Context, automationID string) (map[string]any, *core.Response, error) {
	body := map[string]string{"automation": automationID}
	httpReq, err := s.client.NewRequest(ctx, "POST", ChangeSecretExecuteURL, body)
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

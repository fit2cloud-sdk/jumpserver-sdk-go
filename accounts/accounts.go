package accounts

import (
	"context"

	"github.com/jumpserver-south/jumpserver-sdk-go/internal/core"
	"github.com/jumpserver-south/jumpserver-sdk-go/internal/sdkutil"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

// Account URL constants.
const (
	ListURL   = "/api/v1/accounts/accounts/"
	DetailURL = "/api/v1/accounts/accounts/%s/"
	BulkURL   = "/api/v1/accounts/accounts/bulk/"
	SecretURL = "/api/v1/accounts/account-secrets/%s/"
)

// Account connectivity testing URL constants (v4).
const (
	VerifyURL       = "/api/v1/accounts/accounts/verify/"
	VerifyDetailURL = "/api/v1/accounts/accounts/%s/verify/"
	VerifyTaskURL   = "/api/v1/accounts/accounts/verify/"
)

// Service handles /api/v1/accounts.
type Service struct {
	client core.HTTPClient
}

// NewService creates a new accounts Service.
func NewService(c core.HTTPClient) *Service {
	return &Service{client: c}
}

// List returns a paginated list of accounts.
func (s *Service) List(ctx context.Context, opts *core.ListOptions) ([]model.Account, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(ListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.AccountPage
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

// Get fetches an account by ID.
func (s *Service) Get(ctx context.Context, id string) (*model.Account, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(DetailURL, id), nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.Account
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Create creates an account.
func (s *Service) Create(ctx context.Context, req *model.AccountRequest) (*model.Account, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", ListURL, req)
	if err != nil {
		return nil, nil, err
	}
	var out model.Account
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// CreateBulk adds the same account to many assets in one call.
func (s *Service) CreateBulk(ctx context.Context, reqs []model.AccountRequest) (*core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", BulkURL, reqs)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, httpReq, nil)
}

// CreateBulkByTemplate adds accounts to assets using an account template.
func (s *Service) CreateBulkByTemplate(ctx context.Context, req *model.AccountBulkByTemplateRequest) (*core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", BulkURL, req)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, httpReq, nil)
}

// Update patches an account.
func (s *Service) Update(ctx context.Context, id string, req *model.AccountRequest) (*model.Account, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "PATCH", sdkutil.Spath(DetailURL, id), req)
	if err != nil {
		return nil, nil, err
	}
	var out model.Account
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Delete deletes an account.
func (s *Service) Delete(ctx context.Context, id string) (*core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "DELETE", sdkutil.Spath(DetailURL, id), nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, httpReq, nil)
}

// GetSecret fetches the decrypted account secret.
func (s *Service) GetSecret(ctx context.Context, id string) (*model.Account, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(SecretURL, id), nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.Account
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Verify returns the connectivity verification result for an account (v4).
func (s *Service) Verify(ctx context.Context, id string) (*model.AccountVerifyResult, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(VerifyDetailURL, id), nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.AccountVerifyResult
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// CreateVerifyTask creates a connectivity verification task (v4).
func (s *Service) CreateVerifyTask(ctx context.Context, req *model.AccountVerifyTaskRequest) (*model.AccountVerifyTask, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", VerifyTaskURL, req)
	if err != nil {
		return nil, nil, err
	}
	var out model.AccountVerifyTask
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

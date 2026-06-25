package accounts

import (
	"context"

	"github.com/jumpserver-south/jumpserver-sdk-go/internal/core"
	"github.com/jumpserver-south/jumpserver-sdk-go/internal/sdkutil"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

// Account backup plan URL constants.
const (
	BackupListURL    = "/api/v1/accounts/account-backup-plans/"
	BackupDetailURL  = "/api/v1/accounts/account-backup-plans/%s/"
	BackupExecuteURL = "/api/v1/accounts/account-backup-plan-executions/"
)

// BackupService handles /api/v1/accounts/account-backup-plans.
type BackupService struct {
	client core.HTTPClient
}

// NewBackupService creates a new BackupService.
func NewBackupService(c core.HTTPClient) *BackupService {
	return &BackupService{client: c}
}

// List returns a paginated list of account backup plans.
func (s *BackupService) List(ctx context.Context, opts *core.ListOptions) ([]model.AccountBackupPlan, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(BackupListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.AccountBackupPlanPage
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

// Get fetches an account backup plan by ID.
func (s *BackupService) Get(ctx context.Context, id string) (*model.AccountBackupPlan, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(BackupDetailURL, id), nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.AccountBackupPlan
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Create creates an account backup plan.
func (s *BackupService) Create(ctx context.Context, req *model.AccountBackupPlanRequest) (*model.AccountBackupPlan, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", BackupListURL, req)
	if err != nil {
		return nil, nil, err
	}
	var out model.AccountBackupPlan
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Update patches an account backup plan.
func (s *BackupService) Update(ctx context.Context, id string, req *model.AccountBackupPlanRequest) (*model.AccountBackupPlan, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "PATCH", sdkutil.Spath(BackupDetailURL, id), req)
	if err != nil {
		return nil, nil, err
	}
	var out model.AccountBackupPlan
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Delete deletes an account backup plan.
func (s *BackupService) Delete(ctx context.Context, id string) (*core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "DELETE", sdkutil.Spath(BackupDetailURL, id), nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, httpReq, nil)
}

// Execute triggers a backup plan execution.
func (s *BackupService) Execute(ctx context.Context, planID string) (map[string]any, *core.Response, error) {
	body := map[string]string{"plan": planID}
	httpReq, err := s.client.NewRequest(ctx, "POST", BackupExecuteURL, body)
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

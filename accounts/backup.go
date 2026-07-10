package accounts

import (
	"context"

	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/core"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/util"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/model"
)

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
	return util.List[model.AccountBackupPlan](ctx, s.client, BackupListURL, opts)
}

// Get fetches an account backup plan by ID.
func (s *BackupService) Get(ctx context.Context, id string) (*model.AccountBackupPlan, *core.Response, error) {
	return util.Get[model.AccountBackupPlan](ctx, s.client, BackupDetailURL, id)
}

// Create creates an account backup plan.
func (s *BackupService) Create(ctx context.Context, req *model.AccountBackupPlanRequest) (*model.AccountBackupPlan, *core.Response, error) {
	return util.Create[model.AccountBackupPlan, model.AccountBackupPlanRequest](ctx, s.client, BackupListURL, req)
}

// Update patches an account backup plan.
func (s *BackupService) Update(ctx context.Context, id string, req *model.AccountBackupPlanRequest) (*model.AccountBackupPlan, *core.Response, error) {
	return util.Update[model.AccountBackupPlan, model.AccountBackupPlanRequest](ctx, s.client, BackupDetailURL, id, req)
}

// Delete deletes an account backup plan.
func (s *BackupService) Delete(ctx context.Context, id string) (*core.Response, error) {
	return util.Delete(ctx, s.client, BackupDetailURL, id)
}
// BatchDelete deletes multiple account backup plans by ID using the cache-then-delete pattern.
func (s *BackupService) BatchDelete(ctx context.Context, ids []string) error {
	return util.BatchDelete(ctx, s.client, BackupListURL, ids)
}

// Execute triggers a backup plan execution.
func (s *BackupService) Execute(ctx context.Context, planID string) (map[string]any, *core.Response, error) {
	body := map[string]string{"plan": planID}
	return util.MapAction(ctx, s.client, BackupExecuteURL, body)
}

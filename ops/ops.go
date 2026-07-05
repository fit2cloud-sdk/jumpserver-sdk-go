package ops

import (
	"context"

	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/core"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/util"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/model"
)

const (
	JobListURL            = "/api/v1/ops/jobs/"
	JobExecutionDetailURL = "/api/v1/ops/job-execution/task-detail/%s/"
)

// Service handles /api/v1/ops endpoints.
type Service struct {
	client core.HTTPClient
}

// NewService creates a new ops Service.
func NewService(c core.HTTPClient) *Service {
	return &Service{client: c}
}

// CreateJob creates a quick-command job.
func (s *Service) CreateJob(ctx context.Context, req *model.OpsJobRequest) (*model.OpsJobResponse, *core.Response, error) {
	return util.Create[model.OpsJobResponse, model.OpsJobRequest](ctx, s.client, JobListURL, req)
}

// GetJobResult returns the execution result of a quick-command job.
func (s *Service) GetJobResult(ctx context.Context, taskID string) (*model.OpsJobResult, *core.Response, error) {
	return util.Get[model.OpsJobResult](ctx, s.client, JobExecutionDetailURL, taskID)
}

package audits

import (
	"context"

	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/core"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/util"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/model"
)

// Login / Operate log URL constants.
const (
	LoginLogListURL          = "/api/v1/audits/login-logs/"
	OperateLogListURL        = "/api/v1/audits/operate-logs/"
	PasswordChangeLogListURL = "/api/v1/audits/password-change-logs/"
	JobLogListURL            = "/api/v1/audits/job-logs/"
)

// ListLoginLogs returns a paginated list of login logs.
func (s *Service) ListLoginLogs(ctx context.Context, opts *core.ListOptions) ([]model.LoginLog, *core.Response, error) {
	return util.List[model.LoginLog](ctx, s.client, LoginLogListURL, opts)
}

// ListOperateLogs returns a paginated list of operate logs.
func (s *Service) ListOperateLogs(ctx context.Context, opts *core.ListOptions) ([]model.OperateLog, *core.Response, error) {
	return util.List[model.OperateLog](ctx, s.client, OperateLogListURL, opts)
}

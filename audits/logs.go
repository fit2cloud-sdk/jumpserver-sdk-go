package audits

import (
	"context"

	"github.com/jumpserver-south/jumpserver-sdk-go/internal/core"
	"github.com/jumpserver-south/jumpserver-sdk-go/internal/sdkutil"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
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
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(LoginLogListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.LoginLogPage
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

// ListOperateLogs returns a paginated list of operate logs.
func (s *Service) ListOperateLogs(ctx context.Context, opts *core.ListOptions) ([]model.OperateLog, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(OperateLogListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.OperateLogPage
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

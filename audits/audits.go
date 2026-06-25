package audits

import (
	"context"
	"io"

	"github.com/jumpserver-south/jumpserver-sdk-go/internal/core"
	"github.com/jumpserver-south/jumpserver-sdk-go/internal/sdkutil"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

// Session URL constants.
const (
	SessionListURL   = "/api/v1/terminal/sessions/"
	SessionDetailURL = "/api/v1/terminal/sessions/%s/"
	SessionReplayURL = "/api/v1/terminal/sessions/%s/replay/"
)

// Service handles session, command, FTP, login, and operate log endpoints.
type Service struct {
	client core.HTTPClient
}

// NewService creates a new audits Service.
func NewService(c core.HTTPClient) *Service {
	return &Service{client: c}
}

// ListSessions returns a paginated list of sessions.
func (s *Service) ListSessions(ctx context.Context, opts *core.ListOptions) ([]model.Session, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(SessionListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.SessionPage
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

// GetSession fetches a session by ID.
func (s *Service) GetSession(ctx context.Context, id string) (*model.Session, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(SessionDetailURL, id), nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.Session
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// DownloadReplay streams the session replay archive into w.
func (s *Service) DownloadReplay(ctx context.Context, sessionID string, w io.Writer) (*core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(SessionReplayURL, sessionID), nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "*/*")
	resp, err := s.client.DoRaw(ctx, httpReq, w)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

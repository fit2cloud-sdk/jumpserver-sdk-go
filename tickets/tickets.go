package tickets

import (
	"context"

	"github.com/jumpserver-south/jumpserver-sdk-go/internal/core"
	"github.com/jumpserver-south/jumpserver-sdk-go/internal/sdkutil"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

const (
	ListURL       = "/api/v1/tickets/tickets/"
	DetailURL     = "/api/v1/tickets/tickets/%s/"
	ApproveURL    = "/api/v1/tickets/tickets/%s/approve/"
	FlowListURL   = "/api/v1/tickets/flows/"
	FlowDetailURL = "/api/v1/tickets/flows/%s/"
)

// Service handles /api/v1/tickets.
type Service struct {
	client core.HTTPClient
}

// NewService creates a new tickets Service.
func NewService(c core.HTTPClient) *Service {
	return &Service{client: c}
}

// List returns a paginated list of tickets.
func (s *Service) List(ctx context.Context, opts *core.ListOptions) ([]model.Ticket, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(ListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.TicketPage
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

// Get fetches a ticket by ID.
func (s *Service) Get(ctx context.Context, id string) (*model.Ticket, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(DetailURL, id), nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.Ticket
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Create opens an asset-application ticket.
func (s *Service) Create(ctx context.Context, req *model.TicketRequest) (*model.Ticket, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", ListURL, req)
	if err != nil {
		return nil, nil, err
	}
	var out model.Ticket
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Approve approves a ticket with action "approve" or "reject".
func (s *Service) Approve(ctx context.Context, id, action string) (*core.Response, error) {
	body := map[string]string{"action": action}
	httpReq, err := s.client.NewRequest(ctx, "POST", sdkutil.Spath(ApproveURL, id), body)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, httpReq, nil)
}

// ListFlows returns a paginated list of ticket flows (workflow definitions).
func (s *Service) ListFlows(ctx context.Context, opts *core.ListOptions) ([]model.TicketFlow, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(FlowListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.TicketFlowPage
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

// UpdateFlow patches a ticket flow definition.
func (s *Service) UpdateFlow(ctx context.Context, id string, req *model.TicketFlowRequest) (*model.TicketFlow, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "PATCH", sdkutil.Spath(FlowDetailURL, id), req)
	if err != nil {
		return nil, nil, err
	}
	var out model.TicketFlow
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

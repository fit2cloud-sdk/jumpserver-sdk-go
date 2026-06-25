package assets

import (
	"context"

	"github.com/jumpserver-south/jumpserver-sdk-go/internal/core"
	"github.com/jumpserver-south/jumpserver-sdk-go/internal/sdkutil"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

// Nodes URL constants.
const (
	NodeListURL   = "/api/v1/assets/nodes/"
	NodeDetailURL = "/api/v1/assets/nodes/%s/"
)

// NodesService handles /api/v1/assets/nodes.
type NodesService struct {
	client core.HTTPClient
}

// NewNodesService creates a new NodesService.
func NewNodesService(c core.HTTPClient) *NodesService {
	return &NodesService{client: c}
}

// List returns a paginated list of nodes.
func (s *NodesService) List(ctx context.Context, opts *core.ListOptions) ([]model.Node, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := sdkutil.AppendQuery(NodeListURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page model.NodePage
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

// Get fetches a node by ID.
func (s *NodesService) Get(ctx context.Context, id string) (*model.Node, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", sdkutil.Spath(NodeDetailURL, id), nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.Node
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Create creates a node.
func (s *NodesService) Create(ctx context.Context, req *model.NodeRequest) (*model.Node, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "POST", NodeListURL, req)
	if err != nil {
		return nil, nil, err
	}
	var out model.Node
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Update patches a node.
func (s *NodesService) Update(ctx context.Context, id string, req *model.NodeRequest) (*model.Node, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "PATCH", sdkutil.Spath(NodeDetailURL, id), req)
	if err != nil {
		return nil, nil, err
	}
	var out model.Node
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Delete deletes a node.
func (s *NodesService) Delete(ctx context.Context, id string) (*core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "DELETE", sdkutil.Spath(NodeDetailURL, id), nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, httpReq, nil)
}

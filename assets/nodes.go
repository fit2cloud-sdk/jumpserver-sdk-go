package assets

import (
	"context"

	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/core"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/sdkutil"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/model"
)

// Nodes URL constants.
const (
	NodeListURL          = "/api/v1/assets/nodes/"
	NodeDetailURL        = "/api/v1/assets/nodes/%s/"
	NodeChildrenTreeURL  = "/api/v1/assets/nodes/children/tree/"
	NodeChildrenURL      = "/api/v1/assets/nodes/%s/children/"
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
	return sdkutil.List[model.Node](ctx, s.client, NodeListURL, opts)
}

// Get fetches a node by ID.
func (s *NodesService) Get(ctx context.Context, id string) (*model.Node, *core.Response, error) {
	return sdkutil.Get[model.Node](ctx, s.client, NodeDetailURL, id)
}

// Create creates a node.
func (s *NodesService) Create(ctx context.Context, req *model.NodeRequest) (*model.Node, *core.Response, error) {
	return sdkutil.Create[model.Node, model.NodeRequest](ctx, s.client, NodeListURL, req)
}

// Update patches a node.
func (s *NodesService) Update(ctx context.Context, id string, req *model.NodeRequest) (*model.Node, *core.Response, error) {
	return sdkutil.Update[model.Node, model.NodeRequest](ctx, s.client, NodeDetailURL, id, req)
}

// Delete deletes a node.
func (s *NodesService) Delete(ctx context.Context, id string) (*core.Response, error) {
	return sdkutil.Delete(ctx, s.client, NodeDetailURL, id)
}

// ChildrenTree returns the node children tree list.
func (s *NodesService) ChildrenTree(ctx context.Context, key string) ([]model.NodeTreeItem, *core.Response, error) {
	params := map[string]string{}
	if key != "" {
		params["key"] = key
	}
	path := sdkutil.AppendQuery(NodeChildrenTreeURL, params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var page core.Page[model.NodeTreeItem]
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

// CreateChild creates a child node under the given parent node.
// Returns 400 error if a sibling node with the same name already exists.
func (s *NodesService) CreateChild(ctx context.Context, parentID string, req *model.NodeChildRequest) (*model.NodeTreeItem, *core.Response, error) {
	return sdkutil.Action[model.NodeTreeItem, model.NodeChildRequest](ctx, s.client, sdkutil.Spath(NodeChildrenURL, parentID), req)
}

package users

import (
	"context"
	"fmt"

	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/core"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/util"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/model"
)

const (
	ProfileURL       = "/api/v1/users/profile/"
	ListURL          = "/api/v1/users/users/"
	DetailURL        = "/api/v1/users/users/%s/"
	InviteURL        = "/api/v1/users/users/invite/"
	PreferenceURL    = "/api/v1/users/preference/"
	AssetAccountsURL = "/api/v1/perms/users/%s/assets/%s/"
	PermsAssetsURL   = "/api/v1/perms/users/%s/assets/"
)

// Service handles /api/v1/users endpoints.
type Service struct {
	client core.HTTPClient
}

// NewService creates a new users Service.
func NewService(c core.HTTPClient) *Service {
	return &Service{client: c}
}

// List returns a paginated list of users. Pass nil filters for no
// resource-specific filtering; common pagination goes in opts.
func (s *Service) List(ctx context.Context, filters map[string]string, opts *core.ListOptions) ([]model.User, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	for k, v := range filters {
		if v != "" {
			params[k] = v
		}
	}
	return util.ListWithParams[model.User](ctx, s.client, ListURL, params)
}

// Get fetches a user by ID.
func (s *Service) Get(ctx context.Context, id string) (*model.User, *core.Response, error) {
	return util.Get[model.User](ctx, s.client, DetailURL, id)
}

// Profile fetches the currently-authenticated user.
func (s *Service) Profile(ctx context.Context) (*model.User, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", ProfileURL, nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.User
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Create creates a new user.
func (s *Service) Create(ctx context.Context, req *model.UserRequest) (*model.User, *core.Response, error) {
	return util.Create[model.User, model.UserRequest](ctx, s.client, ListURL, req)
}

// Update patches a user.
func (s *Service) Update(ctx context.Context, id string, req *model.UserRequest) (*model.User, *core.Response, error) {
	return util.Update[model.User, model.UserRequest](ctx, s.client, DetailURL, id, req)
}

// Replace replaces a user.
func (s *Service) Replace(ctx context.Context, id string, req *model.UserRequest) (*model.User, *core.Response, error) {
	return util.Replace[model.User, model.UserRequest](ctx, s.client, DetailURL, id, req)
}

// Delete deletes a user.
func (s *Service) Delete(ctx context.Context, id string) (*core.Response, error) {
	return util.Delete(ctx, s.client, DetailURL, id)
}
// BatchDelete deletes multiple users by ID using the cache-then-delete pattern.
func (s *Service) BatchDelete(ctx context.Context, ids []string) error {
	return util.BatchDelete(ctx, s.client, ListURL, ids)
}

// Invite invites existing users into the current organization.
func (s *Service) Invite(ctx context.Context, userIDs []string, orgRoles []string) (*core.Response, error) {
	body := map[string][]string{"users": userIDs, "org_roles": orgRoles}
	httpReq, err := s.client.NewRequest(ctx, "POST", InviteURL, body)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, httpReq, nil)
}

// ListGroups lists the groups a user belongs to.
func (s *Service) ListGroups(ctx context.Context, userID string, opts *core.ListOptions) ([]model.Group, *core.Response, error) {
	params := map[string]string{}
	if opts != nil {
		opts.Apply(params)
	}
	path := util.AppendQuery(fmt.Sprintf("/api/v1/users/users/%s/groups/", userID), params)
	httpReq, err := s.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}
	var out []model.Group
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return out, resp, nil
}

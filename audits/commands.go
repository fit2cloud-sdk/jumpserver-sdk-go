package audits

import (
	"context"

	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/core"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/sdkutil"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/model"
)

const (
	SessionCommandURL = "/api/v1/terminal/commands/"
)

// ListCommands returns a paginated list of session commands.
func (s *Service) ListCommands(ctx context.Context, opts *core.ListOptions) ([]model.Command, *core.Response, error) {
	return sdkutil.List[model.Command](ctx, s.client, SessionCommandURL, opts)
}

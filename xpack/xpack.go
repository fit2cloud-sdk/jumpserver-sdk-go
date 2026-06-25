package xpack

import (
	"context"

	"github.com/jumpserver-south/jumpserver-sdk-go/internal/core"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

const (
	LicenseDetailURL = "/api/v1/xpack/license/detail"
)

// Service handles Xpack endpoints (license, SSO). Only available on
// enterprise editions.
type Service struct {
	client core.HTTPClient
}

// NewService creates a new xpack Service.
func NewService(c core.HTTPClient) *Service {
	return &Service{client: c}
}

// License returns the Xpack license details.
func (s *Service) License(ctx context.Context) (*model.License, *core.Response, error) {
	httpReq, err := s.client.NewRequest(ctx, "GET", LicenseDetailURL, nil)
	if err != nil {
		return nil, nil, err
	}
	var out model.License
	resp, err := s.client.Do(ctx, httpReq, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

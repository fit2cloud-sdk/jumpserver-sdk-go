package util

import (
	"context"

	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/core"
)

// CacheURL is the API endpoint for caching resource IDs before batch operations.
const CacheURL = "/api/v1/common/resources/cache/"

// CacheResources posts resource IDs to the cache endpoint and returns the SPM token.
// This is used before batch delete operations to obtain a token that references
// the cached set of resources.
func CacheResources(ctx context.Context, client core.HTTPClient, ids []string) (string, *core.Response, error) {
	body := map[string][]string{"resources": ids}
	out := map[string]any{}
	resp, err := MapActionInPlace(ctx, client, CacheURL, body, &out)
	if err != nil {
		return "", resp, err
	}
	spm, _ := out["spm"].(string)
	return spm, resp, nil
}

// BatchDelete caches the given resource IDs and then sends a DELETE request
// to the specified list URL with the returned SPM token.
// This is the idiomatic Go equivalent of the Python pattern:
//
//	spm = self._cache_resources(ids)
//	requests.delete(url, params={"spm": spm})
func BatchDelete(ctx context.Context, client core.HTTPClient, listURL string, ids []string) error {
	spm, _, err := CacheResources(ctx, client, ids)
	if err != nil {
		return err
	}
	path := AppendQuery(listURL, map[string]string{"spm": spm})
	httpReq, err := client.NewRequest(ctx, "DELETE", path, nil)
	if err != nil {
		return err
	}
	_, err = client.Do(ctx, httpReq, nil)
	return err
}

// MapActionInPlace POSTs a body to an action URL and unmarshals the response into out.
func MapActionInPlace(ctx context.Context, client core.HTTPClient, actionURL string, body, out any) (*core.Response, error) {
	httpReq, err := client.NewRequest(ctx, "POST", actionURL, body)
	if err != nil {
		return nil, err
	}
	return client.Do(ctx, httpReq, out)
}

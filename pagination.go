package jumpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/fit2cloud-sdk/jumpserver-sdk-go/internal/core"
)

// Type aliases re-export core types so users can access them from the
// root jumpserver package.

// ListOptions configures pagination and search for list endpoints.
type ListOptions = core.ListOptions

// Response is the typed wrapper returned by every service call.
type Response = core.Response

// PageFetcher paginates through all pages of a list endpoint.
type PageFetcher = core.PageFetcher

// WalkPages repeatedly invokes fetch with an advancing ListOptions
// until no next page is reported. maxPages bounds the total number of
// HTTP calls as a safety valve; pass 0 for "unbounded".
func WalkPages(ctx context.Context, initial *ListOptions, maxPages int, fetch PageFetcher) error {
	if initial == nil {
		initial = &ListOptions{Limit: 100}
	}
	if initial.Limit <= 0 {
		initial.Limit = 100
	}
	opts := *initial
	pages := 0
	for {
		resp, err := fetch(ctx, &opts)
		pages++
		if err != nil {
			return err
		}
		if !resp.HasNextPage() {
			return nil
		}
		if maxPages > 0 && pages >= maxPages {
			return fmt.Errorf("jumpserver: walk exceeded maxPages=%d", maxPages)
		}
		if resp.StatusCode == http.StatusTooManyRequests {
			time.Sleep(parseRetryAfter(resp.Response))
		}
		next := opts.Next()
		if next == nil {
			return nil
		}
		opts = *next
	}
}

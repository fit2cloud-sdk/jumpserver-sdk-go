package main

import (
	"context"

	jumpserver "github.com/jumpserver-south/jumpserver-sdk-go"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

func testWalkPages() {
	section("WalkPages")

	var all []string
	err := jumpserver.WalkPages(ctx, &jumpserver.ListOptions{Limit: 10}, 20,
		func(ctx2 context.Context, opts *jumpserver.ListOptions) (*jumpserver.Response, error) {
			usrs, resp, err := client.Users.List(ctx2, nil, opts)
			if err != nil {
				return resp, err
			}
			for _, u := range usrs {
				all = append(all, u.Username)
			}
			return resp, nil
		})
	if err != nil {
		fail("WalkPages", err)
	} else {
		ok("WalkPages (" + itoa(len(all)) + " users)")
	}

	section("WithOrgScope")
	s2 := client.WithOrgScope(model.JMSDefaultOrg)
	p, _, err := s2.Users.Profile(ctx)
	if err != nil {
		fail("WithOrgScope(ROOT).Profile", err)
	} else {
		ok("WithOrgScope(ROOT).Profile (" + p.Username + ")")
	}
}

package main

import (
	jumpserver "github.com/jumpserver-south/jumpserver-sdk-go"
)

func testPlatforms() {
	section("Platforms")

	platforms, _, err := client.Platforms.List(ctx, &jumpserver.ListOptions{Limit: 15})
	if err != nil {
		fail("Platforms.List", err)
	} else {
		ok("Platforms.List (" + itoa(len(platforms)) + ")")
	}
	if len(platforms) > 0 {
		p, _, err := client.Platforms.Get(ctx, platforms[0].ID)
		if err != nil {
			fail("Platforms.Get", err)
		} else {
			ok("Platforms.Get (" + p.Name + ")")
		}
	}
}

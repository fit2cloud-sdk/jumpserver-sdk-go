package main

import (
	jumpserver "github.com/jumpserver-south/jumpserver-sdk-go"
)

func testRoles() {
	section("Roles")

	orgRoles, _, err := client.Roles.List(ctx, "org", &jumpserver.ListOptions{Limit: 15})
	if err != nil {
		fail("Roles.List(org)", err)
	} else {
		ok("Roles.List(org) (" + itoa(len(orgRoles)) + ")")
	}

	sysRoles, _, err := client.Roles.List(ctx, "system", &jumpserver.ListOptions{Limit: 15})
	if err != nil {
		fail("Roles.List(system)", err)
	} else {
		ok("Roles.List(system) (" + itoa(len(sysRoles)) + ")")
	}
	if len(orgRoles) > 0 {
		r, _, err := client.Roles.Get(ctx, "org", orgRoles[0].ID)
		if err != nil {
			fail("Roles.Get(org)", err)
		} else {
			ok("Roles.Get(org) (" + r.Name + ")")
		}
	}
}

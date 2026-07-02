package main

import (
	jumpserver "github.com/fit2cloud-sdk/jumpserver-sdk-go"
)

func testTickets() {
	section("Tickets")

	list, _, err := client.Tickets.List(ctx, &jumpserver.ListOptions{Limit: 15})
	if err != nil {
		fail("Tickets.List", err)
	} else {
		ok("Tickets.List (" + itoa(len(list)) + ")")
	}
	if len(list) > 0 {
		t, _, err := client.Tickets.Get(ctx, list[0].ID)
		if err != nil {
			fail("Tickets.Get", err)
		} else {
			ok("Tickets.Get (" + t.Title + ")")
		}
	}
	flows, _, err := client.Tickets.ListFlows(ctx, &jumpserver.ListOptions{Limit: 15})
	if err != nil {
		fail("Tickets.ListFlows", err)
	} else {
		ok("Tickets.ListFlows (" + itoa(len(flows)) + ")")
	}
}

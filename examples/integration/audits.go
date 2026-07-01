package main

import (
	jumpserver "github.com/jumpserver-south/jumpserver-sdk-go"
)

func testAudits() {
	section("Audits")

	sessions, _, err := client.Audits.ListSessions(ctx, &jumpserver.ListOptions{Limit: 3})
	if err != nil {
		fail("Audits.ListSessions", err)
	} else {
		ok("Audits.ListSessions (" + itoa(len(sessions)) + ")")
	}
	if len(sessions) > 0 {
		s, _, err := client.Audits.GetSession(ctx, sessions[0].ID)
		if err != nil {
			fail("Audits.GetSession", err)
		} else {
			ok("Audits.GetSession (" + s.Asset + ")")
		}
	}
	cmds, _, err := client.Audits.ListCommands(ctx, &jumpserver.ListOptions{Limit: 3})
	if err != nil {
		fail("Audits.ListCommands", err)
	} else {
		ok("Audits.ListCommands (" + itoa(len(cmds)) + ")")
	}
	ftp, _, err := client.Audits.ListFTPLogs(ctx, &jumpserver.ListOptions{Limit: 3})
	if err != nil {
		fail("Audits.ListFTPLogs", err)
	} else {
		ok("Audits.ListFTPLogs (" + itoa(len(ftp)) + ")")
	}
	ll, _, err := client.Audits.ListLoginLogs(ctx, &jumpserver.ListOptions{Limit: 3})
	if err != nil {
		fail("Audits.ListLoginLogs", err)
	} else {
		ok("Audits.ListLoginLogs (" + itoa(len(ll)) + ")")
	}
	ol, _, err := client.Audits.ListOperateLogs(ctx, &jumpserver.ListOptions{Limit: 3})
	if err != nil {
		fail("Audits.ListOperateLogs", err)
	} else {
		ok("Audits.ListOperateLogs (" + itoa(len(ol)) + ")")
	}
}

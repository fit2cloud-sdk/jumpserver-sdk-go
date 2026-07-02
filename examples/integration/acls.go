package main

import (
	jumpserver "github.com/fit2cloud-sdk/jumpserver-sdk-go"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/model"
)

func testCommandFilters() {
	section("CommandFilters CRUD")

	cf, _, err := scoped.CommandFilters.Create(ctx, &model.CommandFilterRequest{
		Name: "cf-" + ts, Action: "reject",
		Users: map[string]any{"type": "all"}, Assets: map[string]any{"type": "all"},
		Accounts: []string{"*"},
	})
	if err != nil {
		fail("CommandFilters.Create", err)
	} else {
		createdCmdFilterID = cf.ID
		ok("CommandFilters.Create (id=" + cf.ID + ")")
		u, _, err := scoped.CommandFilters.Update(ctx, cf.ID, &model.CommandFilterRequest{
			Name: "cf-upd-" + ts, Action: "reject",
			Users: map[string]any{"type": "all"}, Assets: map[string]any{"type": "all"},
			Accounts: []string{"*"},
		})
		if err != nil {
			fail("CommandFilters.Update", err)
		} else {
			ok("CommandFilters.Update (" + u.Name + ")")
		}
		got, _, err := scoped.CommandFilters.Get(ctx, cf.ID)
		if err != nil {
			fail("CommandFilters.Get", err)
		} else {
			ok("CommandFilters.Get (" + got.Name + ")")
		}
		list, _, err := scoped.CommandFilters.List(ctx, &jumpserver.ListOptions{Limit: 15})
		if err != nil {
			fail("CommandFilters.List", err)
		} else {
			ok("CommandFilters.List (" + itoa(len(list)) + ")")
		}
	}
}

func testCommandGroups() {
	section("CommandGroups CRUD")

	cg, _, err := scoped.CommandFilters.CreateGroup(ctx, &model.CommandGroupRequest{
		Name: "cg-" + ts, Type: map[string]string{"value": "command"}, Content: "rm -rf",
	})
	if err != nil {
		fail("CommandGroups.Create", err)
	} else {
		createdCmdGroupID = cg.ID
		ok("CommandGroups.Create (id=" + cg.ID + ")")
		u, _, err := scoped.CommandFilters.UpdateGroup(ctx, cg.ID, &model.CommandGroupRequest{
			Name: "cg-upd-" + ts, Type: map[string]string{"value": "command"}, Content: "rm -rf /",
		})
		if err != nil {
			fail("CommandGroups.Update", err)
		} else {
			ok("CommandGroups.Update (" + u.Name + ")")
		}
		got, _, err := scoped.CommandFilters.GetGroup(ctx, cg.ID)
		if err != nil {
			fail("CommandGroups.Get", err)
		} else {
			ok("CommandGroups.Get (" + got.Name + ")")
		}
		list, _, err := scoped.CommandFilters.ListGroups(ctx, &jumpserver.ListOptions{Limit: 15})
		if err != nil {
			fail("CommandGroups.List", err)
		} else {
			ok("CommandGroups.List (" + itoa(len(list)) + ")")
		}
	}
}

func testLoginACLs() {
	section("LoginACLs")

	list, _, err := client.LoginACLs.List(ctx, &jumpserver.ListOptions{Limit: 15})
	if err != nil {
		fail("LoginACLs.List", err)
	} else {
		ok("LoginACLs.List (" + itoa(len(list)) + ")")
	}
	if len(list) > 0 {
		a, _, err := client.LoginACLs.Get(ctx, list[0].ID)
		if err != nil {
			fail("LoginACLs.Get", err)
		} else {
			ok("LoginACLs.Get (" + a.Name + ")")
		}
	}
}

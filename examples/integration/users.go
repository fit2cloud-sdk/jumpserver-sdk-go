package main

import (
	jumpserver "github.com/jumpserver-south/jumpserver-sdk-go"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

func testUsers() {
	section("Users")

	profile, _, err := client.Users.Profile(ctx)
	if err != nil {
		fail("Users.Profile", err)
	} else {
		ok("Users.Profile (" + profile.Username + ")")
	}

	users, _, err := client.Users.List(ctx, nil, &jumpserver.ListOptions{Limit: 15})
	if err != nil {
		fail("Users.List", err)
	} else {
		ok("Users.List (" + itoa(len(users)) + ")")
	}
	if len(users) > 0 {
		u, _, err := client.Users.Get(ctx, users[0].ID)
		if err != nil {
			fail("Users.Get", err)
		} else {
			ok("Users.Get (" + u.Username + ")")
		}
	}

	section("UserGroups CRUD")

	g, _, err := scoped.UserGroups.Create(ctx, &model.GroupRequest{Name: "ug-" + ts})
	if err != nil {
		fail("UserGroups.Create", err)
	} else {
		ok("UserGroups.Create (id=" + g.ID + ")")
		u, _, err := scoped.UserGroups.Update(ctx, g.ID, &model.GroupRequest{Name: "ug-upd-" + ts})
		if err != nil {
			fail("UserGroups.Update", err)
		} else {
			ok("UserGroups.Update (" + u.Name + ")")
		}
		got, _, err := scoped.UserGroups.Get(ctx, g.ID)
		if err != nil {
			fail("UserGroups.Get", err)
		} else {
			ok("UserGroups.Get (" + got.Name + ")")
		}
		list, _, err := scoped.UserGroups.List(ctx, &jumpserver.ListOptions{Limit: 15})
		if err != nil {
			fail("UserGroups.List", err)
		} else {
			ok("UserGroups.List (" + itoa(len(list)) + ")")
		}
		lu, _, err := scoped.UserGroups.ListUsers(ctx, g.ID, &jumpserver.ListOptions{Limit: 15})
		if err != nil {
			skip("UserGroups.ListUsers", err.Error())
		} else {
			ok("UserGroups.ListUsers (" + itoa(len(lu)) + ")")
		}
		_, err = scoped.UserGroups.Delete(ctx, g.ID)
		if err != nil {
			fail("UserGroups.Delete", err)
		} else {
			ok("UserGroups.Delete")
		}
	}
}

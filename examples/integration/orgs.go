package main

import (
	jumpserver "github.com/fit2cloud-sdk/jumpserver-sdk-go"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/model"
)

func testOrgs() {
	section("Organizations CRUD")

	o, _, err := client.Organizations.Create(ctx, &model.OrganizationRequest{Name: "org-" + ts})
	if err != nil {
		fail("Organizations.Create", err)
	} else {
		ok("Organizations.Create (id=" + o.ID + ")")
		u, _, err := client.Organizations.Update(ctx, o.ID, &model.OrganizationRequest{Name: "org-upd-" + ts})
		if err != nil {
			fail("Organizations.Update", err)
		} else {
			ok("Organizations.Update (" + u.Name + ")")
		}
		got, _, err := client.Organizations.Get(ctx, o.ID)
		if err != nil {
			fail("Organizations.Get", err)
		} else {
			ok("Organizations.Get (" + got.Name + ")")
		}
		list, _, err := client.Organizations.List(ctx, &jumpserver.ListOptions{Limit: 15})
		if err != nil {
			fail("Organizations.List", err)
		} else {
			ok("Organizations.List (" + itoa(len(list)) + ")")
		}
		_, err = client.Organizations.Delete(ctx, o.ID)
		if err != nil {
			fail("Organizations.Delete", err)
		} else {
			ok("Organizations.Delete")
		}
	}
}

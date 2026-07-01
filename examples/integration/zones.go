package main

import (
	jumpserver "github.com/jumpserver-south/jumpserver-sdk-go"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

func testZones() {
	section("Zones CRUD")

	z, _, err := scoped.Zones.Create(ctx, &model.ZoneRequest{Name: "zone-" + ts})
	if err != nil {
		fail("Zones.Create", err)
	} else {
		createdZoneID = z.ID
		ok("Zones.Create (id=" + z.ID + ")")
		u, _, err := scoped.Zones.Update(ctx, z.ID, &model.ZoneRequest{Name: "zone-upd-" + ts})
		if err != nil {
			fail("Zones.Update", err)
		} else {
			ok("Zones.Update (" + u.Name + ")")
		}
		got, _, err := scoped.Zones.Get(ctx, z.ID)
		if err != nil {
			fail("Zones.Get", err)
		} else {
			ok("Zones.Get (" + got.Name + ")")
		}
		list, _, err := scoped.Zones.List(ctx, &jumpserver.ListOptions{Limit: 15})
		if err != nil {
			fail("Zones.List", err)
		} else {
			ok("Zones.List (" + itoa(len(list)) + ")")
		}
	}
}

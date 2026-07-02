package main

import (
	jumpserver "github.com/fit2cloud-sdk/jumpserver-sdk-go"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/model"
)

func testLabels() {
	section("Labels CRUD")

	l, _, err := scoped.Labels.Create(ctx, &model.LabelRequest{Name: "lbl-" + ts, Value: "v1"})
	if err != nil {
		fail("Labels.Create", err)
	} else {
		createdLabelID = l.ID
		ok("Labels.Create (id=" + l.ID + ")")
		u, _, err := scoped.Labels.Update(ctx, l.ID, &model.LabelRequest{Name: "lbl-upd-" + ts, Value: "v2"})
		if err != nil {
			fail("Labels.Update", err)
		} else {
			ok("Labels.Update (" + u.Name + ")")
		}
		got, _, err := scoped.Labels.Get(ctx, l.ID)
		if err != nil {
			fail("Labels.Get", err)
		} else {
			ok("Labels.Get (" + got.Name + ")")
		}
		list, _, err := scoped.Labels.List(ctx, &jumpserver.ListOptions{Limit: 15})
		if err != nil {
			fail("Labels.List", err)
		} else {
			ok("Labels.List (" + itoa(len(list)) + ")")
		}
	}
}

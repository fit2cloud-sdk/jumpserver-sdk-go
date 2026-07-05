package main

import (
	jumpserver "github.com/fit2cloud-sdk/jumpserver-sdk-go"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/model"
)

func testNodes() {
	section("Nodes CRUD")

	n, _, err := scoped.Nodes.Create(ctx, &model.NodeRequest{Value: "node-" + ts})
	if err != nil {
		fail("Nodes.Create", err)
	} else {
		createdNodeID = n.ID
		ok("Nodes.Create (id=" + n.ID + ")")
		u, _, err := scoped.Nodes.Update(ctx, n.ID, &model.NodeRequest{Value: "node-upd-" + ts})
		if err != nil {
			fail("Nodes.Update", err)
		} else {
			ok("Nodes.Update (" + u.Value + ")")
		}
		got, _, err := scoped.Nodes.Get(ctx, n.ID)
		if err != nil {
			fail("Nodes.Get", err)
		} else {
			ok("Nodes.Get (" + got.Value + ")")
		}
		list, _, err := scoped.Nodes.List(ctx, &jumpserver.ListOptions{Limit: 15})
		if err != nil {
			fail("Nodes.List", err)
		} else {
			ok("Nodes.List (" + itoa(len(list)) + ")")
		}
	}

	section("Node Children Tree")
	tree, _, err := scoped.Nodes.ChildrenTree(ctx, "")
	if err != nil {
		fail("Nodes.ChildrenTree", err)
	} else {
		ok("Nodes.ChildrenTree (" + itoa(len(tree)) + " nodes)")
		if len(tree) > 0 {
			child, _, err := scoped.Nodes.CreateChild(ctx, tree[0].Meta.Data.ID, &model.NodeChildRequest{
				Value: "child-" + ts,
			})
			if err != nil {
				fail("Nodes.CreateChild", err)
			} else {
				ok("Nodes.CreateChild (value=" + child.Value + ")")
			}
		}
	}
}

package main

import (
	"fmt"

	jumpserver "github.com/fit2cloud-sdk/jumpserver-sdk-go"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/model"
)

func testPermissions() {
	section("Permissions CRUD")

	p, _, err := scoped.Permissions.Create(ctx, &model.AssetPermissionRequest{Name: "perm-" + ts})
	if err != nil {
		fail("Permissions.Create", err)
	} else {
		createdPermID = p.ID
		ok("Permissions.Create (id=" + p.ID + ")")
		u, _, err := scoped.Permissions.Update(ctx, p.ID, &model.AssetPermissionRequest{Name: "perm-upd-" + ts})
		if err != nil {
			fail("Permissions.Update", err)
		} else {
			ok("Permissions.Update (" + u.Name + ")")
		}
		got, _, err := scoped.Permissions.Get(ctx, p.ID)
		if err != nil {
			fail("Permissions.Get", err)
		} else {
			ok("Permissions.Get (" + got.Name + ")")
		}
		list, _, err := scoped.Permissions.List(ctx, &jumpserver.ListOptions{Limit: 15})
		if err != nil {
			fail("Permissions.List", err)
		} else {
			ok("Permissions.List (" + itoa(len(list)) + ")")
		}
		_, err = scoped.Permissions.Delete(ctx, p.ID)
		if err != nil {
			fail("Permissions.Delete", err)
		} else {
			ok("Permissions.Delete")
			createdPermID = ""
		}
	}
}

func testSelfAssets() {
	section("Self Assets")

	list, _, err := client.Self.ListAssets(ctx, nil, &jumpserver.ListOptions{Limit: 15})
	if err != nil {
		fail("Self.ListAssets", err)
	} else {
		ok("Self.ListAssets (" + itoa(len(list)) + ")")
		if len(list) > 0 {
			detail, _, err := client.Self.GetAsset(ctx, list[0].ID)
			if err != nil {
				fail("Self.GetAsset", err)
			} else {
				ok("Self.GetAsset (" + detail.Name + ")")
			}
		}
	}
}

func testPermRelations() {
	section("Permission Relations")

	perm, _, err := scoped.Permissions.Create(ctx, &model.AssetPermissionRequest{Name: "perm-rel-" + ts})
	if err != nil {
		skip("PermRelations", "create perm failed: "+err.Error())
		return
	}
	defer scoped.Permissions.Delete(ctx, perm.ID)

	profile, _, err := client.Users.Profile(ctx)
	if err != nil {
		skip("PermRelations", "profile failed: "+err.Error())
		return
	}

	users, _, err := client.Permissions.AddUsersRelations(ctx, []model.AssetPermUserRelation{
		{User: profile.ID, AssetPermission: perm.ID},
	})
	if err != nil {
		fail("PermRelations.AddUsers", err)
	} else {
		ok("PermRelations.AddUsers (" + itoa(len(users)) + ")")
	}

	assetList, _, err := scoped.Hosts.List(ctx, nil, &jumpserver.ListOptions{Limit: 1})
	if err == nil && len(assetList) > 0 {
		ar, _, err := client.Permissions.AddAssetsRelations(ctx, []model.AssetPermAssetRelation{
			{Asset: assetList[0].ID, AssetPermission: perm.ID},
		})
		if err != nil {
			fail("PermRelations.AddAssets", err)
		} else {
			ok("PermRelations.AddAssets (" + itoa(len(ar)) + ")")
		}
	} else {
		skip("PermRelations.AddAssets", "no assets available")
	}

	list, _, err := scoped.Nodes.List(ctx, &jumpserver.ListOptions{Limit: 1})
	if err == nil && len(list) > 0 {
		nr, _, err := client.Permissions.AddNodesRelations(ctx, []model.AssetPermNodeRelation{
			{Node: list[0].ID, AssetPermission: perm.ID},
		})
		if err != nil {
			fail("PermRelations.AddNodes", err)
		} else {
			ok("PermRelations.AddNodes (" + itoa(len(nr)) + ")")
		}
	} else {
		skip("PermRelations.AddNodes", "no nodes available")
	}

	groups, _, err := scoped.UserGroups.List(ctx, &jumpserver.ListOptions{Limit: 1})
	if err == nil && len(groups) > 0 {
		gr, _, err := client.Permissions.AddUserGroupsRelations(ctx, []model.AssetPermUserGroupRelation{
			{UserGroup: groups[0].ID, AssetPermission: perm.ID},
		})
		if err != nil {
			fail("PermRelations.AddUserGroups", err)
		} else {
			ok("PermRelations.AddUserGroups (" + itoa(len(gr)) + ")")
		}
	} else {
		skip("PermRelations.AddUserGroups", "no user groups available")
	}

	// Verify that the perm has these relations by getting its detail
	got, _, err := scoped.Permissions.Get(ctx, perm.ID)
	if err != nil {
		fail("PermRelations.Verify", err)
	} else if got != nil {
		ok(fmt.Sprintf("PermRelations.Verify (perm=%s, users=%d, assets=%d, nodes=%d, groups=%d)",
			got.Name, len(got.Users), len(got.Assets), len(got.Nodes), len(got.UserGroups)))
	}
}

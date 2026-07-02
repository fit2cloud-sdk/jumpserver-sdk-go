package main

import (
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

	profile, _, err := client.Users.Profile(ctx)
	if err != nil {
		skip("PermRelations", "profile failed: "+err.Error())
		scoped.Permissions.Delete(ctx, perm.ID)
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

	assets, _, err := client.Permissions.AddAssetsRelations(ctx, []model.AssetPermAssetRelation{
		{Asset: "00000000-0000-0000-0000-000000000000", AssetPermission: perm.ID},
	})
	if err != nil {
		fail("PermRelations.AddAssets", err)
	} else {
		ok("PermRelations.AddAssets (" + itoa(len(assets)) + ")")
	}

	nodes, _, err := client.Permissions.AddNodesRelations(ctx, []model.AssetPermNodeRelation{
		{Node: "00000000-0000-0000-0000-000000000000", AssetPermission: perm.ID},
	})
	if err != nil {
		fail("PermRelations.AddNodes", err)
	} else {
		ok("PermRelations.AddNodes (" + itoa(len(nodes)) + ")")
	}

	groups, _, err := client.Permissions.AddUserGroupsRelations(ctx, []model.AssetPermUserGroupRelation{
		{UserGroup: "00000000-0000-0000-0000-000000000000", AssetPermission: perm.ID},
	})
	if err != nil {
		fail("PermRelations.AddUserGroups", err)
	} else {
		ok("PermRelations.AddUserGroups (" + itoa(len(groups)) + ")")
	}

	_, err = scoped.Permissions.Delete(ctx, perm.ID)
	if err != nil {
		fail("PermRelations.Cleanup", err)
	}
}

package main

import (
	"fmt"
	"strings"

	jumpserver "github.com/fit2cloud-sdk/jumpserver-sdk-go"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/assets"
	"github.com/fit2cloud-sdk/jumpserver-sdk-go/model"
)

func testAssets() {
	testHosts()
	testCategoryAssets()
	testGenericAssets()
	testGateways()
}

func testHosts() {
	section("Hosts CRUD")

	platforms, _, _ := client.Platforms.List(ctx, &jumpserver.ListOptions{Limit: 100})
	var platID int
	for _, p := range platforms {
		if strings.Contains(strings.ToLower(p.Name), "linux") {
			platID = p.ID
			break
		}
	}
	if platID == 0 && len(platforms) > 0 {
		platID = platforms[0].ID
	}
	if platID == 0 {
		skip("Hosts.Create", "no platform")
		return
	}

	h, _, err := scoped.Hosts.Create(ctx, &model.AssetRequest{
		Name: "host-" + ts, Address: "192.168.1." + ts[:min(3, len(ts))],
		Platform: platID, Protocols: []model.NamePort{{Name: "ssh", Port: 22}},
	})
	if err != nil {
		fail("Hosts.Create", err)
		return
	}
	createdHostID = h.ID
	ok("Hosts.Create (id=" + h.ID + ")")

	u, _, err := scoped.Hosts.Update(ctx, h.ID, &model.AssetRequest{
		Name: "host-upd-" + ts, Address: "10.0.0.1", Platform: platID,
		Protocols: []model.NamePort{{Name: "ssh", Port: 2222}},
	})
	if err != nil {
		fail("Hosts.Update", err)
	} else {
		ok("Hosts.Update (" + u.Name + ")")
	}
	got, _, err := scoped.Hosts.Get(ctx, h.ID)
	if err != nil {
		fail("Hosts.Get", err)
	} else {
		ok("Hosts.Get (" + got.Name + ")")
	}
	list, _, err := scoped.Hosts.List(ctx, nil, &jumpserver.ListOptions{Limit: 15})
	if err != nil {
		fail("Hosts.List", err)
	} else {
		ok("Hosts.List (" + itoa(len(list)) + ")")
	}
}

func testCategoryAssets() {
	section("Category assets CRUD (Databases, Devices, Webs, Clouds, Customs)")

	allPlatforms, _, err := client.Platforms.List(ctx, &jumpserver.ListOptions{Limit: 500})
	if err != nil {
		fail("Platforms.List (for categories)", err)
		return
	}
	platByCategory := make(map[string]model.Platform)
	for _, p := range allPlatforms {
		cat := p.Category.Value
		if cat != "" {
			if _, exists := platByCategory[cat]; !exists {
				platByCategory[cat] = p
			}
		}
	}

	type catTest struct {
		name    string
		catKey  string
		svc     *assets.CategoryService
		address string
	}
	cats := []catTest{
		{"Databases", "database", scoped.Databases, "10.0.1.1"},
		{"Devices", "device", scoped.Devices, "10.0.2.1"},
		{"Webs", "web", scoped.Webs, "10.0.3.1"},
		{"Clouds", "cloud", scoped.Clouds, "https://cloud.example.com"},
		{"Customs", "custom", scoped.Customs, "10.0.5.1"},
	}
	for _, tc := range cats {
		p, found := platByCategory[tc.catKey]
		if !found {
			skip(tc.name+".Create", "no platform for category "+tc.catKey)
			continue
		}
		var protocols []model.NamePort
		for _, proto := range p.Protocols {
			if proto.Name != "" {
				protocols = append(protocols, model.NamePort{Name: proto.Name, Port: proto.Port})
			}
		}
		if len(protocols) == 0 {
			protocols = []model.NamePort{{Name: "ssh", Port: 22}}
		}

		asset, _, err := tc.svc.Create(ctx, &model.AssetRequest{
			Name:      fmt.Sprintf("%s-%s", strings.ToLower(tc.name), ts),
			Address:   tc.address,
			Platform:  p.ID,
			Protocols: protocols,
		})
		if err != nil {
			fail(tc.name+".Create", err)
			continue
		}
		createdCategoryIDs[tc.name] = asset.ID
		ok(fmt.Sprintf("%s.Create (id=%s, platform=%s/%d)", tc.name, asset.ID, p.Name, p.ID))

		u, _, err := tc.svc.Update(ctx, asset.ID, &model.AssetRequest{
			Name:      fmt.Sprintf("%s-upd-%s", strings.ToLower(tc.name), ts),
			Address:   tc.address,
			Platform:  p.ID,
			Protocols: protocols,
		})
		if err != nil {
			fail(tc.name+".Update", err)
		} else {
			ok(fmt.Sprintf("%s.Update (%s)", tc.name, u.Name))
		}
		got, _, err := tc.svc.Get(ctx, asset.ID)
		if err != nil {
			fail(tc.name+".Get", err)
		} else {
			ok(fmt.Sprintf("%s.Get (%s)", tc.name, got.Name))
		}
		list, _, err := tc.svc.List(ctx, nil, &jumpserver.ListOptions{Limit: 15})
		if err != nil {
			fail(tc.name+".List", err)
		} else {
			ok(fmt.Sprintf("%s.List (%d)", tc.name, len(list)))
		}
	}
}

func testGenericAssets() {
	section("Assets (generic)")

	assetList, _, err := client.Assets.List(ctx, nil, &jumpserver.ListOptions{Limit: 15})
	if err != nil {
		fail("Assets.List", err)
	} else {
		ok("Assets.List (" + itoa(len(assetList)) + ")")
	}
	if len(assetList) > 0 {
		a, _, err := client.Assets.Get(ctx, assetList[0].ID)
		if err != nil {
			fail("Assets.Get", err)
		} else {
			ok("Assets.Get (" + a.Name + ")")
		}
	}
}

func testGateways() {
	section("Gateways CRUD")

	platforms, _, _ := client.Platforms.List(ctx, &jumpserver.ListOptions{Limit: 200})
	var gwPlat int
	for _, p := range platforms {
		if strings.Contains(strings.ToLower(p.Name), "gateway") || strings.Contains(strings.ToLower(p.Name), "网关") {
			gwPlat = p.ID
			break
		}
	}
	if gwPlat == 0 {
		skip("Gateways.Create", "no gateway-type platform found")
		return
	}

	gw, _, err := scoped.Gateways.Create(ctx, &model.GatewayRequest{
		Name: "gw-" + ts, Address: "172.16.0.1", Platform: gwPlat,
		Protocols: []model.NamePort{{Name: "ssh", Port: 22}},
	})
	if err != nil {
		fail("Gateways.Create", err)
	} else {
		ok("Gateways.Create (id=" + gw.ID + ")")
		u, _, err := scoped.Gateways.Update(ctx, gw.ID, &model.GatewayRequest{
			Name: "gw-upd-" + ts, Address: "172.16.0.2", Platform: gwPlat,
			Protocols: []model.NamePort{{Name: "ssh", Port: 22}},
		})
		if err != nil {
			fail("Gateways.Update", err)
		} else {
			ok("Gateways.Update (" + u.Name + ")")
		}
		got, _, err := scoped.Gateways.Get(ctx, gw.ID)
		if err != nil {
			fail("Gateways.Get", err)
		} else {
			ok("Gateways.Get (" + got.Name + ")")
		}
		_, err = scoped.Gateways.Delete(ctx, gw.ID)
		if err != nil {
			fail("Gateways.Delete", err)
		} else {
			ok("Gateways.Delete")
		}
	}
	list, _, err := scoped.Gateways.List(ctx, &jumpserver.ListOptions{Limit: 15})
	if err != nil {
		fail("Gateways.List", err)
	} else {
		ok("Gateways.List (" + itoa(len(list)) + ")")
	}
}

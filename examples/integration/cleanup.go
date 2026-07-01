package main

func cleanup() {
	section("Cleanup")

	for _, c := range []struct {
		name string
		fn   func() (bool, error)
	}{
		{"Account", func() (bool, error) {
			if createdAccountID == "" {
				return true, nil
			}
			_, e := scoped.Accounts.Delete(ctx, createdAccountID)
			return false, e
		}},
		{"AccountTemplate", func() (bool, error) {
			if createdTemplateID == "" {
				return true, nil
			}
			_, e := scoped.AccountTemplates.Delete(ctx, createdTemplateID)
			return false, e
		}},
		{"Host", func() (bool, error) {
			if createdHostID == "" {
				return true, nil
			}
			_, e := scoped.Hosts.Delete(ctx, createdHostID)
			return false, e
		}},
		{"Category:Database", func() (bool, error) {
			id := createdCategoryIDs["Databases"]
			if id == "" {
				return true, nil
			}
			_, e := scoped.Databases.Delete(ctx, id)
			return false, e
		}},
		{"Category:Device", func() (bool, error) {
			id := createdCategoryIDs["Devices"]
			if id == "" {
				return true, nil
			}
			_, e := scoped.Devices.Delete(ctx, id)
			return false, e
		}},
		{"Category:Web", func() (bool, error) {
			id := createdCategoryIDs["Webs"]
			if id == "" {
				return true, nil
			}
			_, e := scoped.Webs.Delete(ctx, id)
			return false, e
		}},
		{"Category:Cloud", func() (bool, error) {
			id := createdCategoryIDs["Clouds"]
			if id == "" {
				return true, nil
			}
			_, e := scoped.Clouds.Delete(ctx, id)
			return false, e
		}},
		{"Category:Custom", func() (bool, error) {
			id := createdCategoryIDs["Customs"]
			if id == "" {
				return true, nil
			}
			_, e := scoped.Customs.Delete(ctx, id)
			return false, e
		}},
		{"CommandGroup", func() (bool, error) {
			if createdCmdGroupID == "" {
				return true, nil
			}
			_, e := scoped.CommandFilters.DeleteGroup(ctx, createdCmdGroupID)
			return false, e
		}},
		{"CommandFilter", func() (bool, error) {
			if createdCmdFilterID == "" {
				return true, nil
			}
			_, e := scoped.CommandFilters.Delete(ctx, createdCmdFilterID)
			return false, e
		}},
		{"Permission", func() (bool, error) {
			if createdPermID == "" {
				return true, nil
			}
			_, e := scoped.Permissions.Delete(ctx, createdPermID)
			return false, e
		}},
		{"Node", func() (bool, error) {
			if createdNodeID == "" {
				return true, nil
			}
			_, e := scoped.Nodes.Delete(ctx, createdNodeID)
			return false, e
		}},
		{"Zone", func() (bool, error) {
			if createdZoneID == "" {
				return true, nil
			}
			_, e := scoped.Zones.Delete(ctx, createdZoneID)
			return false, e
		}},
		{"Label", func() (bool, error) {
			if createdLabelID == "" {
				return true, nil
			}
			_, e := scoped.Labels.Delete(ctx, createdLabelID)
			return false, e
		}},
	} {
		skipped, err := c.fn()
		if err != nil {
			fail("Cleanup."+c.name, err)
		} else if skipped {
			skip("Cleanup."+c.name, "not created")
		} else {
			ok("Cleanup." + c.name)
		}
	}
}

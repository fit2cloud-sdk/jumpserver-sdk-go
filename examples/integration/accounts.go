package main

import (
	jumpserver "github.com/jumpserver-south/jumpserver-sdk-go"
	"github.com/jumpserver-south/jumpserver-sdk-go/model"
)

func testAccountTemplates() {
	section("AccountTemplates CRUD")

	t, _, err := scoped.AccountTemplates.Create(ctx, &model.AccountTemplateRequest{
		Name: "tpl-" + ts, Username: "testuser", SecretType: "password",
	})
	if err != nil {
		fail("AccountTemplates.Create", err)
	} else {
		createdTemplateID = t.ID
		ok("AccountTemplates.Create (id=" + t.ID + ")")
		u, _, err := scoped.AccountTemplates.Update(ctx, t.ID, &model.AccountTemplateRequest{
			Name: "tpl-upd-" + ts, Username: "testuser2", SecretType: "password",
		})
		if err != nil {
			fail("AccountTemplates.Update", err)
		} else {
			ok("AccountTemplates.Update (" + u.Name + ")")
		}
		got, _, err := scoped.AccountTemplates.Get(ctx, t.ID)
		if err != nil {
			fail("AccountTemplates.Get", err)
		} else {
			ok("AccountTemplates.Get (" + got.Name + ")")
		}
		list, _, err := scoped.AccountTemplates.List(ctx, &jumpserver.ListOptions{Limit: 15})
		if err != nil {
			fail("AccountTemplates.List", err)
		} else {
			ok("AccountTemplates.List (" + itoa(len(list)) + ")")
		}
	}
}

func testAccounts() {
	section("Accounts CRUD")

	if createdHostID == "" {
		skip("Accounts.Create", "no host")
		return
	}

	a, _, err := scoped.Accounts.Create(ctx, &model.AccountRequest{
		Name: "acct-" + ts, Username: "testacct", Asset: createdHostID,
		SecretType: "password", Secret: "TestPass123!",
	})
	if err != nil {
		fail("Accounts.Create", err)
	} else {
		createdAccountID = a.ID
		ok("Accounts.Create (id=" + a.ID + ")")
		u, _, err := scoped.Accounts.Update(ctx, a.ID, &model.AccountRequest{
			Username: "testacct-upd", Asset: createdHostID, SecretType: "password",
		})
		if err != nil {
			fail("Accounts.Update", err)
		} else {
			ok("Accounts.Update (" + u.Username + ")")
		}
		got, _, err := scoped.Accounts.Get(ctx, a.ID)
		if err != nil {
			fail("Accounts.Get", err)
		} else {
			ok("Accounts.Get (" + got.Username + ")")
		}
		list, _, err := scoped.Accounts.List(ctx, &jumpserver.ListOptions{Limit: 15})
		if err != nil {
			fail("Accounts.List", err)
		} else {
			ok("Accounts.List (" + itoa(len(list)) + ")")
		}
	}
}

func testChangeSecrets() {
	section("ChangeSecrets CRUD")

	cs, _, err := scoped.ChangeSecrets.Create(ctx, &model.ChangeSecretAutomationRequest{
		Name: "cs-" + ts, Accounts: []string{"root"}, SecretType: "password",
		SecretStrategy: "specific", IsPeriodic: true, Interval: 24,
	})
	if err != nil {
		fail("ChangeSecrets.Create", err)
	} else {
		ok("ChangeSecrets.Create (id=" + cs.ID + ")")
		u, _, err := scoped.ChangeSecrets.Update(ctx, cs.ID, &model.ChangeSecretAutomationRequest{
			Name: "cs-upd-" + ts, Accounts: []string{"root"}, SecretType: "password",
			SecretStrategy: "specific", IsPeriodic: true, Interval: 48,
		})
		if err != nil {
			fail("ChangeSecrets.Update", err)
		} else {
			ok("ChangeSecrets.Update (" + u.Name + ")")
		}
		got, _, err := scoped.ChangeSecrets.Get(ctx, cs.ID)
		if err != nil {
			fail("ChangeSecrets.Get", err)
		} else {
			ok("ChangeSecrets.Get (" + got.Name + ")")
		}
		list, _, err := scoped.ChangeSecrets.List(ctx, &jumpserver.ListOptions{Limit: 15})
		if err != nil {
			fail("ChangeSecrets.List", err)
		} else {
			ok("ChangeSecrets.List (" + itoa(len(list)) + ")")
		}
		_, err = scoped.ChangeSecrets.Delete(ctx, cs.ID)
		if err != nil {
			fail("ChangeSecrets.Delete", err)
		} else {
			ok("ChangeSecrets.Delete")
		}
	}
}

func testAccountBackups() {
	section("AccountBackups CRUD")

	bp, _, err := scoped.AccountBackups.Create(ctx, &model.AccountBackupPlanRequest{
		Name: "bp-" + ts, Accounts: []string{"root"}, SecretType: "password",
		IsPeriodic: true, Interval: 24,
	})
	if err != nil {
		fail("AccountBackups.Create", err)
	} else {
		ok("AccountBackups.Create (id=" + bp.ID + ")")
		u, _, err := scoped.AccountBackups.Update(ctx, bp.ID, &model.AccountBackupPlanRequest{
			Name: "bp-upd-" + ts, Accounts: []string{"root"}, SecretType: "password",
			IsPeriodic: true, Interval: 48,
		})
		if err != nil {
			fail("AccountBackups.Update", err)
		} else {
			ok("AccountBackups.Update (" + u.Name + ")")
		}
		got, _, err := scoped.AccountBackups.Get(ctx, bp.ID)
		if err != nil {
			fail("AccountBackups.Get", err)
		} else {
			ok("AccountBackups.Get (" + got.Name + ")")
		}
		list, _, err := scoped.AccountBackups.List(ctx, &jumpserver.ListOptions{Limit: 15})
		if err != nil {
			fail("AccountBackups.List", err)
		} else {
			ok("AccountBackups.List (" + itoa(len(list)) + ")")
		}
		_, err = scoped.AccountBackups.Delete(ctx, bp.ID)
		if err != nil {
			fail("AccountBackups.Delete", err)
		} else {
			ok("AccountBackups.Delete")
		}
	}
}

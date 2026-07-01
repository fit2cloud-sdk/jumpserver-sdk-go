package main

import "fmt"

func testSettings() {
	section("Settings")

	pub, _, err := client.Settings.Public(ctx)
	if err != nil {
		fail("Settings.Public", err)
	} else {
		ok("Settings.Public (watermark=" + boolStr(pub.EnableWatermark) + ")")
	}

	settings, _, err := client.Settings.List(ctx, nil)
	if err != nil {
		fail("Settings.List", err)
	} else {
		ok("Settings.List (" + itoa(len(settings)) + " keys)")
	}
}

func boolStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func itoa(n int) string {
	return fmt.Sprintf("%d", n)
}

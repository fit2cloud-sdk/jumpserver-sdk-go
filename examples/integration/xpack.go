package main

func testXpack() {
	section("Xpack")

	lic, _, err := client.Xpack.License(ctx)
	if err != nil {
		fail("Xpack.License", err)
	} else {
		ok("Xpack.License (valid=" + boolStr(lic.IsValid) + ", edition=" + lic.Edition + ")")
	}
}

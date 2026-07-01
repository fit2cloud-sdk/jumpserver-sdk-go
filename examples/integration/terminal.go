package main

func testTerminal() {
	section("Terminal")

	methods, _, err := client.Terminal.ConnectMethods(ctx)
	if err != nil {
		fail("Terminal.ConnectMethods", err)
	} else {
		ok("Terminal.ConnectMethods (" + itoa(len(methods)) + " keys)")
	}
}

package main

import (
	"os"
	"testing"

	mullvad "./AmIMullvad"
)

var sampleData mullvad.MullvadResult

// TestMain --
// setup method for JSON parsing tests.
func TestMain(m *testing.M) {
	sampleData = mullvad.ParseSample()
	os.Exit(m.Run())
}

// TestIP --
// Make sure the IP was properly parsed.
func TestIP(t *testing.T) {
	if sampleData.IP != "127.0.0.1" {
		t.Fail()
	}
}

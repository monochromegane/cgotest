package cgotest

import "testing"

func TestGetHome(t *testing.T) {
	if GetHome() == "nothing" {
		t.Errorf("CGO is disable.")
	}
}

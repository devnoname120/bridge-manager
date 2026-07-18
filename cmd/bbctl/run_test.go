package main

import "testing"

func TestLocalBuildScript(t *testing.T) {
	tests := map[string]string{
		"instagram": "./build-ig.sh",
		"meta":      "./build.sh",
		"telegram":  "./build.sh",
	}
	for bridgeType, expected := range tests {
		if actual := localBuildScript(bridgeType); actual != expected {
			t.Errorf("localBuildScript(%q) = %q, expected %q", bridgeType, actual, expected)
		}
	}
}

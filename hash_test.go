package main

import "testing"

func TestGetHash(t *testing.T) {
	exampleURL := "https://www.google.com"
	expectedHash := "8739bc55"
	exampleHash := getHash(exampleURL)
	if expectedHash != exampleHash {
		t.Errorf("unexpected hash '%s', expected '%s'", exampleHash, expectedHash)
	}
}

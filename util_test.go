package main

import "testing"

func TestValidURL(t *testing.T) {
	url := "https://www.google.com"
	if !isValidProtocol(url) {
		t.Errorf("URL '%s' was marked as an invalid url, expected: true", url)
	}
}

func TestNonValidURL(t *testing.T) {
	url := "ftp://www.google.com"
	if isValidProtocol(url) {
		t.Errorf("URL '%s' was marked as valid, expected: false", url)
	}
}

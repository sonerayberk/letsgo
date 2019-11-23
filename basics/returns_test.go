package main

import "testing"

func TestReturns(t *testing.T) {
	got := returns()
	want := "func returns a string"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

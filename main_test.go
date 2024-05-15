package main

import "testing"

func TestCap(t *testing.T) {
	expected := "John"
	result := Cap("john")
	if result != expected {
		t.Errorf("got %s expected %s", result, expected)
		t.Errorf("Try again")
	}
}
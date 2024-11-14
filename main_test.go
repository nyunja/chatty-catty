package main

import (
	"chatty/core"
	"testing"
)

func TestCap(t *testing.T) {
	expected := "John"
	result := core.Cap("john")
	if result != expected {
		t.Errorf("got %s expected %s", result, expected)
		t.Errorf("Try again")
	}
}

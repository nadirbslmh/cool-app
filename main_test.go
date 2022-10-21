package main

import "testing"

func TestAdd(t *testing.T) {
	var expected int = 11

	var actual int = Add(5, 7)

	if expected != actual {
		t.Errorf("expected: %d got: %d", expected, actual)
	}
}

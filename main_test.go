package main_test

import "testing"

func TestAddition(t *testing.T) {
	got := 2 + 2
	expected := 4
	if got != expected {
		t.Errorf("Wrong result. Got: '%v', wanted: '%v'", got, expected)
	}
}

func TestSubtraction(t *testing.T) {
	got := 10 - 6
	expected := 4
	if got != expected {
		t.Errorf("Wrong result. Got: '%v', wanted: '%v'", got, expected)
	}
}

package main

import "testing"

var neg = "0123456789a"
var pos = "bcdef"

func TestIsOpenNegative(t *testing.T) {
	for i := range neg {
		if isOpen(neg[i]) {
			t.Errorf("Expected %v got %v for %v", false, true, string(neg[i]))
		}
	}
}

func TestIsOpenPositive(t *testing.T) {
	for i := range pos {
		if !isOpen(pos[i]) {
			t.Errorf("Expected %v got %v for %v", false, true, string(neg[i]))
		}
	}
}

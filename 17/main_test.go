package main

import "testing"

var negative = "0123456789a"
var positive = "bcdef"

func TestIsOpenNegative(t *testing.T) {
	for i := range negative {
		if isOpen(negative[i]) {
			t.Errorf("Expected %v got %v for %v", false, true, string(negative[i]))
		}
	}
}

func TestIsOpenPositive(t *testing.T) {
	for i := range positive {
		if !isOpen(positive[i]) {
			t.Errorf("Expected %v got %v for %v", true, false, string(positive[i]))
		}
	}
}

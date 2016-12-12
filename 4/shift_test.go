package main

import "testing"

func TestWholePhrase(t *testing.T) {
	res := shift("qzmt-zixmtkozy-ivhz", 343)
	if res != "very encrypted name" {
		t.Errorf("Expected %v got %v", "very encrypted name", res)
	}
}

func TestSingleLetter(t *testing.T) {
	for l := 'a'; l < 'y'; l++ {
		res := shift(string(l), 1)
		if res != string(l+1) {
			t.Errorf("Expected %v got %v", string(l+1), res)
		}
	}
	res := shift(string('z'), 1)
	if res != "a" {
		t.Errorf("Expected %v got %v", "a", res)
	}
}

func TestWrap(t *testing.T) {
	res := shift(string('z'), 1)
	if res != "a" {
		t.Errorf("Expected %v got %v", "a", res)
	}
}

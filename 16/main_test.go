package main

import "testing"

var checkCases = []testcase{testcase{"110010110100", "100"}, testcase{"10000011110010000111", "01100"}}

func TestChecksum(t *testing.T) {
	for _, c := range checkCases {
		res := checksum(*initialize(c.input, len(c.input)))
		if res != c.expected {
			t.Errorf("Expected %v got %v", c.expected, res)
		}
	}
}

type testcase struct {
	input, expected string
}

var fillCases = []testcase{testcase{"1", "100"}, testcase{"0", "001"},
	testcase{"11111", "11111000000"}, testcase{"111100001010", "1111000010100101011110000"},
	testcase{"10000", "10000011110"},
	testcase{"10000011110", "10000011110010000111110"}}

func TestFilling1Step(t *testing.T) {
	for _, c := range fillCases {
		limit := len(c.input)*2 + 1
		tot := initialize(c.input, limit)
		fillDisk(tot, limit, len(c.input))
		if print(*tot) != c.expected {
			t.Errorf("Expected %v got %v", c.expected, print(*tot))
		}
	}
}

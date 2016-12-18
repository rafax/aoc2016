package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var res string
	if os.Getenv("PART") != "1" {
		res = part1("01110110101001000", 35651584)
	} else {
		res = part2("ffykfhsq")
	}
	fmt.Println(res)
}

func part1(prefix string, limit int) string {
	tot := initialize(prefix)
	tot = fillDisk(tot, limit)
	return checksum(tot)
}
func part2(prefix string) string { return "" }

func initialize(prefix string) []bool {
	tot := make([]bool, len(prefix))
	for i, l := range prefix {
		if l == '1' {
			tot[i] = true
		}
	}
	// fmt.Printf("Initialized %v to %v\n", prefix, tot)
	return tot
}

func fillDisk(tot []bool, limit int) []bool {
	for len(tot) < limit {
		rev := make([]bool, len(tot))
		for i := 0; i < len(tot); i++ {
			rev[len(tot)-1-i] = !tot[i]
		}
		tot = append(append(tot, false), rev...)
	}
	return tot[:limit]
}

func checksum(in []bool) string {
	for len(in)%2 == 0 {
		for i := 0; i < len(in)-1; i += 2 {
			if in[i] == in[i+1] {
				in[i/2] = true
			} else {
				in[i/2] = false
			}
		}
		in = in[:len(in)/2]
	}
	return print(in)
}

func print(in []bool) string {
	acc := make([]string, len(in))
	for i, v := range in {
		if v {
			acc[i] = "1"
		} else {
			acc[i] = "0"
		}
	}
	return strings.Join(acc, "")
}

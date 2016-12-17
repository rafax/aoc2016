package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var res string
	if os.Getenv("PART") != "1" {
		res = part1("10000", 20)
	} else {
		res = part2("ffykfhsq")
	}
	fmt.Println(res)
}

func part1(prefix string, limit int) string {
	tot := initialize(prefix, limit)
	fillDisk(tot, limit, len(prefix))
	return checksum(*tot)
}
func part2(prefix string) string { return "" }

func initialize(prefix string, limit int) *[]bool {
	tot := make([]bool, limit)
	for i, l := range prefix {
		if l == '1' {
			tot[i] = true
		}
	}
	return &tot
}

func fillDisk(tot *[]bool, limit, init int) {
	//fmt.Println("Starting fill with ", *tot, limit, init)
	offset := init
	width := init
	for {
		for i := 0; i < width; i++ {
			//	fmt.Printf("Step(%v) with %v: l:%v off:%v w:%v\n", i, *tot, limit, offset, width)
			if offset+i+1 >= limit {
				return
			}
			(*tot)[offset+i+1] = !(*tot)[offset-1-i]
		}
		width = 2*width + 1
	}
}

func checksum(in []bool) string {
	limit := len(in)
	for limit%2 == 0 {
		for i := 0; i < len(in)-1; i += 2 {
			if in[i] == in[i+1] {
				in[i/2] = true
			} else {
				in[i/2] = false
			}
		}
		limit /= 2
	}
	return print(in[:limit])
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

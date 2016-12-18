package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var res string
	if os.Getenv("PART") == "1" {
		res = part1(".^..^....^....^^.^^.^.^^.^.....^.^..^...^^^^^^.^^^^.^.^^^^^^^.^^^^^..^.^^^.^^..^.^^.^....^.^...^^.^.", 40)
	} else {
		res = part1(".^..^....^....^^.^^.^.^^.^.....^.^..^...^^^^^^.^^^^.^.^^^^^^^.^^^^^..^.^^^.^^..^.^^.^....^.^...^^.^.", 400000)
	}
	fmt.Println(res)
}

func parse(in string) ([]bool, int) {
	safe := 0
	res := make([]bool, len(in))
	for i, v := range in {
		if v == '^' {
			res[i] = true
		} else {
			safe++
		}
	}
	return res, safe
}

func part1(in string, steps int) string {
	row, safe := parse(in)
	for i := 0; i < steps-1; i++ {
		new := make([]bool, len(row))
		for i, v := range row {
			l, c, r := false, v, false
			if i > 0 {
				l = row[i-1]
			}
			if i < len(row)-1 {
				r = row[i+1]
			}
			if (l && c && !r) || (!l && c && r) || (l && !c && !r) || (!l && !c && r) {
				new[i] = true
			} else {
				safe++
			}
		}
		row = new
	}
	fmt.Println(safe)
	return print(row)
}

func print(in []bool) string {
	res := make([]string, len(in))
	for i, v := range in {
		if v {
			res[i] = "^"
		} else {
			res[i] = "."
		}
	}
	return strings.Join(res, "")
}
